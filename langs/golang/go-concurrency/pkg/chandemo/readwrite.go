package chandemo

import "fmt"

// 经典的生产者、消费者的多协程工作
// 通道作为一等公民

func worker(id int, c chan int) {
	// 读取chan
	for n := range c {
		fmt.Printf("Worker %d receive %c\n", id, n)
	}
}

// 创建工作协程
func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func Working() {
	var channels [3] chan int

	// 启动并发协程
	for i := 0; i < 3; i++ {
		channels[i] = createWorker(i)
	}

	// 生产者，触发睡眠的工作协程开始工作
	for i := 0; i < 3; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 3; i++ {
		channels[i] <- 'A' + i
	}
}
