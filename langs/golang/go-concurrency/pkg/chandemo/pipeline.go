package chandemo

import "fmt"

// 通道的链式
func Pipeline() {
	natural := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 11; x++ {
			natural <- x
		}
		close(natural)
	}()

	go func() {
		// for..range在通道关闭后，会自动退出，无需检查ok
		for x := range natural {
			squares <- x * x
		}
		close(squares)
	}()

	for y:= range squares{
		fmt.Println(y)
	}
}
