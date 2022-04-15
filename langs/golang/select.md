# select

作者：知乎用户
链接：https://www.zhihu.com/question/441064352/answer/1697603147
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



## 无阻塞获取值

select-default模式，节选自[fasthttp1.19/client.go#L1955-L1963](https://link.zhihu.com/?target=https%3A//github.com/valyala/fasthttp/blob/v1.19.0/client.go%23L1955-L1963)。

```go
// waiting reports whether w is still waiting for an answer (connection or error).
func (w *wantConn) waiting() bool {
	select {
	case <-w.ready:
		return false
	default:
		return true
	}
}
```

## 超时控制

select-timer模式，例如等待tcp节点发送连接包，超时后则关闭连接。

```go
func (n *node) waitForConnectPkt() {
	select {
	case <-n.connected:
		log.Println("接收到连接包")
	case <-time.After(time.Second * 5):
		log.Println("接收连接包超时")
		n.conn.Close()
	}
}
```

## 类事件驱动循环

for-select模式，例如监控tcp节点心跳是否正常。

```go
func (n *node) heartbeatDetect() {
	for {
		select {
		case <-n.heartbeat:
			// 收到心跳信号则退出select等待下一次心跳
			break
		case <-time.After(time.Second*3):
			// 心跳超时，关闭连接
			n.conn.Close()
			return
		}
	}
}
```

## 带优先级的任务队列

嵌套select模式，节选自[k8s1.20.2/taint_manager.go#L244-L275](https://link.zhihu.com/?target=https%3A//github.com/kubernetes/kubernetes/blob/v1.20.2/pkg/controller/nodelifecycle/scheduler/taint_manager.go%23L244-L275)。

```go
func (tc *NoExecuteTaintManager) worker(worker int, done func(), stopCh <-chan struct{}) {
	defer done()

	// When processing events we want to prioritize Node updates over Pod updates,
	// as NodeUpdates that interest NoExecuteTaintManager should be handled as soon as possible -
	// we don't want user (or system) to wait until PodUpdate queue is drained before it can
	// start evicting Pods from tainted Nodes.
	for {
		select {
		case <-stopCh:
			return
		case nodeUpdate := <-tc.nodeUpdateChannels[worker]:
			tc.handleNodeUpdate(nodeUpdate)
			tc.nodeUpdateQueue.Done(nodeUpdate)
		case podUpdate := <-tc.podUpdateChannels[worker]:
			// If we found a Pod update we need to empty Node queue first.
		priority:
			for {
				select {
				case nodeUpdate := <-tc.nodeUpdateChannels[worker]:
					tc.handleNodeUpdate(nodeUpdate)
					tc.nodeUpdateQueue.Done(nodeUpdate)
				default:
					break priority
				}
			}
			// After Node queue is emptied we process podUpdate.
			tc.handlePodUpdate(podUpdate)
			tc.podUpdateQueue.Done(podUpdate)
		}
	}
}
```



> golang 的 select 本质上是展开成 if - else 的形式。本质上就是一个多条件判断。select 可以最经典的可以结合 channel 来使用。select 结合到channel的时候经常作为一种类似io复用的方式。

- https://www.zhihu.com/question/441064352