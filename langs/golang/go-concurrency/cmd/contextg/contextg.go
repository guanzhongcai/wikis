package contextg

import (
	"errors"
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

var Canceled = errors.New("context canceled")
var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct {
}

func (deadlineExceededError) Error() string {
	return "context deadline exceeded"
}

func (deadlineExceededError) Timeout() bool {
	return true
}

func (deadlineExceededError) Temporary() bool {
	return true
}

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknown empty context"
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

type CancelFunc func()

// WithCancel 返回一个parent的copy，且带有一个新的Done channel
// 当返回的cancel函数被调用时，返回的context的Done channel会是关闭的
// 或当parent context的Done channel是关闭的（无论谁先发生）
//
// 取消这个context会释放和它相关的资源，所以代码应该立即调用cancel
// 当操作运行到context完成时
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() {
		c.cancel(true, Canceled)
	}
}

// newCancelCtx 返回初始化了测 cancelCtx
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}

// propagateCancel 安排child取消，当parent是取消
func propagateCancel(parent Context, child canceler) {
	if parent.Done() == nil {
		return // parent从不取消
	}
	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent已经被取消
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
				// nothing
			}
		}()
	}
}

type canceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     chan struct{}         // created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first call
}

// 需要实现Context的各个接口，否则不能在func中多态使用
func (c *cancelCtx) Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

type stringer interface {
	String() string
}

func contextName(c Context) string {
	if s, ok := c.(stringer); ok {
		return s.String()
	}
	// fixme: return reflectlite.TypeOf(c).String()
	return ""
}

func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}
	c.err = err
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}

	for child := range c.children {
		// note: acquiring the child's lock while holding parent's lock
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		removeChild(c.Context, c)
	}
}

// WithDeadline 返回一个parent的copy，带上了调整后不迟于d的参数
// 如果parent的deadline已经早于d，WithDeadline(parent, d)是和parent的语义是等同的
// 返回的context是Done channel是关闭的：
// 1.如果deadline过期了
// 2.如果返回的cancel函数被调用
// 3.parent context的Done channel被关闭
//
// 取消这个context会释放和它有关的资源，同上
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		// 当前的deadline已经快于新的
		return WithCancel(parent)
	}

	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
	propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
		// deadline已经过去了
		c.cancel(true, DeadlineExceeded)
		return c, func() {
			c.cancel(false, Canceled)
		}
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true, DeadlineExceeded)
		})
	}
	return c, func() {
		c.cancel(true, Canceled)
	}
}

// parentCancelCtx 查找parent链的引用，条件：*chancelCtx.
// 此函数知道此包的每个实体类型代表的parent
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	for {
		switch c := parent.(type) {
		case *cancelCtx:
			return c, true
		case *timerCtx:
			return &c.cancelCtx, true
		case *valueCtx:
			parent = c.Context

		default:
			return nil, false
		}
	}
}

// removeChild 从它的parent中移除context
func removeChild(parent Context, child canceler) {
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children, child)
	}
	p.mu.Unlock()
}

type timerCtx struct {
	cancelCtx
	timer    *time.Timer // under cancelCtx.mu
	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *timerCtx) cancel(removeFromParent bool, err error) {
	c.cancelCtx.cancel(false, err)
	if removeFromParent {
		// 从它的parent的cancelCtx's的children中移除这个timerCtx
		removeChild(c.cancelCtx.Context, c)
	}
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}

// WithTimeout 返回WithDeadline的实现
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithValue(parent Context, key, val interface{}) Context {
	if key == nil {
		panic("nil key")
	}
	// todo: key is not comparable
	return &valueCtx{parent, key, val}
}

type valueCtx struct {
	Context
	key, val interface{}
}

// stringify 没有使用fmt包，因为不想context依赖unicode表
func stringify(v interface{}) string {
	switch s := v.(type) {
	case stringer:
		return s.String()
	case string:
		return s
	}
	return "not <Stringer>"
}

func (c *valueCtx) String() string {
	return contextName(c.Context) + ".WithValue(type " +
		stringify(c.key) + ", val " + stringify(c.val) + ")"
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}
