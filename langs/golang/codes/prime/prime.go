/*
要找出10000以内所有的素数，这里使用的方法是筛法，即从2开始每找到一个素数就标记所有能被该素数整除的所有数。直到没有可标记的数，剩下的就都是素数。下面以找出10以内所有素数为例，借用 CSP 方式解决这个问题。
通过4个并发处理程序得出10以内的素数表
*/

package main

import "fmt"

func Processor(seq chan int, wait chan struct{}, id int) {
	id++
	go func(idx int) {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Printf("协程: %d, 素数是: %d\n", idx, prime)
		out := make(chan int)
		Processor(out, wait, idx)		
		for num := range seq {
			if num % prime != 0 {
				fmt.Printf("协程: %d, 标记的数是: %d\n", idx, num)
				out <- num
			}				
		}
		close(out)
	}(id)
}

func main() {
	origin, wait := make(chan int), make(chan struct{})
	id := 0 
	Processor(origin, wait, id)
	// 从2开始每找到一个素数就标记所有能被该素数整除的所有数
	for num := 2; num < 10; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}

/*
output:
协程: 1, 素数是: 2
协程: 1, 标记的数是: 3
协程: 1, 标记的数是: 5
协程: 1, 标记的数是: 7
协程: 1, 标记的数是: 9
协程: 2, 素数是: 3
协程: 2, 标记的数是: 5
协程: 2, 标记的数是: 7
协程: 3, 素数是: 5
协程: 3, 标记的数是: 7
协程: 4, 素数是: 7
*/