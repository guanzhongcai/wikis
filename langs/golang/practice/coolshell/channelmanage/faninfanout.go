package channelmanage

import (
	"fmt"
	"math"
	"sync"
)

// 动用Go语言的 Go Routine和 Channel还有一个好处，就是可以写出1对多，或多对1的pipeline，也就是Fan In/ Fan Out。
// 下面，我们来看一个Fan in的示例：
// 我们想通过并发的方式来对一个很长的数组中的质数进行求和运算，我们想先把数组分段求和，然后再把其集中起来。

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func merge(cs []<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func SumPrimes() {
	nums := makeRange(1, 10000)
	in := echo(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum(prime(in))
	}

	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
}
