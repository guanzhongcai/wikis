package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestPreventGoroutineLeak(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	u := User{Name:"leo", Age:7}
	fmt.Printf("%#v\n", u)

	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {

		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// do sth
					fmt.Println(s)
				case <-done:
					return
				default:
					//fmt.Println("waiting..")
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	s := make(chan string, 1)
	s <- "nil"
	terminated := doWork(done, nil)

	go func() {
		// cancel op after 1 second
		time.Sleep(1 * time.Second)
		fmt.Println("canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done")
}

/** output
canceling doWork goroutine...
doWork exited.
Done
 */
