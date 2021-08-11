package chandemo

// 20个缓冲区的信号量
var semaMax = make(chan struct{}, 20)

func MaxDemo() {

	// 哪怕开的协程再过，同时并发数是固定的，20个！
	semaMax <- struct{}{}
	defer func() {
		<-semaMax
	}()
}
