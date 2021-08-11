package chandemo

import (
	"fmt"
	"time"
)

// 定时的tick通道
func TickDemo() {
	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)

	var nfiles, nbytes int64
	fileSizes := make(chan int64)
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size

		case <-tick: // 每隔500ms打印一次
			fmt.Println("ticking")
		}
	}
}
