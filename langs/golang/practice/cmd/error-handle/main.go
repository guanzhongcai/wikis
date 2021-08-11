// 把结果和错误都回传回来
package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {

	// 返回可读取的channel，以检索循环迭代的结果
	checkStatus := func(done <-chan interface{}, urls ... string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				// 创建一个Result实例，并设置错误和响应字段
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				// 将结果写入channel
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	errCount := 0
	urls := []string{"https://www.baidu.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			errCount++
			if errCount >= 3 {
				// 3个或更多错误的时候停止尝试检查错误
				fmt.Println("too many errors, breaking")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
