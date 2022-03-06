package main

import "math"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		for value%i == 0 {
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

func main() {
	nums := makeRange(1, 10000)
	in := echo(nums)
	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = prime(in)
	}
}
