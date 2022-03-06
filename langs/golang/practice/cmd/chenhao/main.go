package main

import (
	"fmt"
	"github.com/jackpal/gateway"
	natpmp "github.com/jackpal/go-nat-pmp"
	"sync"
)

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}

func (s *Square) Area() int {
	return 0
}

func main() {
	var s Shape
	s = &Square{len: 5}
	fmt.Printf("%d\n", s.Sides())

	pmp()
	range2D()
	utf8range()
}

func pmp() {
	gatewayIP, err := gateway.DiscoverGateway()
	if err != nil {
		fmt.Println("err DiscoverGateway: " + err.Error())
		return
	}

	client := natpmp.NewClient(gatewayIP)
	response, err := client.GetExternalAddress()
	if err != nil {
		fmt.Println("err GetExternalAddress: " + err.Error())
		return
	}
	fmt.Printf("External IP address: %v\n", response.ExternalIPAddress)
}
func utf8range() {
	mu := new(sync.Mutex)
	mu.Lock()
	words := []rune("hello世界")
	for _, word := range words {
		fmt.Println(word)
	}
	mu.Unlock()
}

func range2D() {
	var times [5][0]int
	for range times { // 循环了5次，却不用付出额外的内存代价！
		fmt.Println("hello")
	}
}
