package go_generation

import (
	"fmt"
	"testing"
	"time"
)

func TestUsage(t *testing.T) {
	Usage()
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(2*time.Second)
	<- timer.C
	fmt.Println("timer expired!")
}

func TestTicker(t *testing.T)  {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}

}