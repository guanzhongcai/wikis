// 并发写文件
package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}

	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

// 匿名结构体
func structAnonymous()  {
	var user struct{
		Name string
		Age int
	}
	user.Name = "gzc"
	user.Age = 22
	fmt.Printf("%+v", user)
}

// 生成数的平方 in为只写 out为只读
func generateSquares(in chan <- int, out <- chan int)  {
	for x := range out {
		in <- x * x
	}
	close(in)
	//close(out) // 报错，因为out是只读，此处不可关闭；关闭是写操作；关闭后还可以读，但不可以写
}

func main() {
	structAnonymous()
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}

	go consume(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()

	d := <- done

	if d == true {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
