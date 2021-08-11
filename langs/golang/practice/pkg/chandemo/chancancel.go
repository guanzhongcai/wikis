package chandemo

import (
	"flag"
	"os"
)

// 并发的退出，创建一个chan，触发后续的并发退出

var done = make(chan struct{})
// 创建一个信号量
var sema = make(chan struct{}, 20)

var verbose = flag.Bool("v", false, "show verboses progress messages")

func cancelled() bool {
	select {
	case <-done: // 如果通道关闭，就会进入此case。通过close(done)来实现通知
		return true
	default: // 如果通道未关闭，则进入此case，即未取消，可继续操作
		return false
	}
}

func DemoMain() {

	// 启动一个协程监听任意字符的输入，来关闭通道，以触发并行的<-done的读取事件
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// do your main work
	go walkDir("~/wiki")
	go dirents("~/wiki")
}

func walkDir(dir string) {
	if cancelled() {
		return
	}
	// do sth
}

func dirents(dir string) []os.FileInfo {
	// 任意信号都退出执行（并行检测）
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}

	return nil
}
