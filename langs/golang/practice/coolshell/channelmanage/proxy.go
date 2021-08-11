package channelmanage

import "fmt"

// 使用一个代理函数来完成
type EchoFunc func([]int) (<-chan int)
type PipeFunc func(<-chan int)(<-chan int)

func pipeline(nums []int, echo EchoFunc, pipeFunc ... PipeFunc) <- chan int {

	ch := echo(nums)
	for i := range pipeFunc {
		ch = pipeFunc[i](ch)
	}
	return ch
}

func PipelineExec()  {
	var nums  = []int{1,2,3,4,5,6,7,89,10}
	for n := range pipeline(nums, echo, odd, sq, sum){
		fmt.Println(n)
	}
}
