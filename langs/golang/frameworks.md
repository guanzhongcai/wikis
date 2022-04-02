# Go开源框架

## gin

- [gin日志打印到某个目录，并实现拆分](https://blog.csdn.net/u010918487/article/details/86146691)

- 自定义中间件

  ```go
  func Logger() gin.HandlerFunc {
  	return func(c *gin.Context) {
  		t := time.Now()
  
  		// Set example variable
  		c.Set("example", "12345")
  
  		// before request
  
  		c.Next()
  
  		// after request
  		latency := time.Since(t)
  		log.Print(latency)
  
  		// access the status we are sending
  		status := c.Writer.Status()
  		log.Println(status)
  	}
  }
  
  func main() {
  	r := gin.New()
  	r.Use(Logger())
  
  	r.GET("/test", func(c *gin.Context) {
  		example := c.MustGet("example").(string)
  
  		// it would print: "12345"
  		log.Println(example)
  	})
  
  	// Listen and serve on 0.0.0.0:8080
  	r.Run(":8080")
  }
  ```

- 当在中间件或处理程序中启动新的Goroutines时，您不应该使用它内部的原始上下文，您必须使用只读副本。

  ```go
  func main() {
  	r := gin.Default()
  
  	r.GET("/long_async", func(c *gin.Context) {
  		// create copy to be used inside the goroutine
  		cCp := c.Copy()
  		go func() {
  			// simulate a long task with time.Sleep(). 5 seconds
  			time.Sleep(5 * time.Second)
  
  			// note that you are using the copied context "cCp", IMPORTANT
  			log.Println("Done! in path " + cCp.Request.URL.Path)
  		}()
  	})
  	r.Run(":8080")
  }
  ```

- [graceful stop or restart](https://gin-gonic.com/docs/examples/graceful-restart-or-stop/)

  - http.Server’s built-in [Shutdown()](https://golang.org/pkg/net/http/#Server.Shutdown) method for graceful shutdowns
  - use [fvbock/endless](https://github.com/fvbock/endless) to replace the default `ListenAndServe`
  - [grace](https://github.com/facebookgo/grace): Graceful restart & zero downtime deploy for Go servers.

- HTTP2 server push

  ```go
  	r.GET("/", func(c *gin.Context) {
  		if pusher := c.Writer.Pusher(); pusher != nil {
  			// use pusher.Push() to do server push
  			if err := pusher.Push("/assets/app.js", nil); err != nil {
  				log.Printf("Failed to push: %v", err)
  			}
  		}
  		c.HTML(200, "https", gin.H{
  			"status": "success",
  		})
  	})
  ```



## 参考资料

- [Gin Web Framework中文文档](https://gin-gonic.com/zh-cn/docs/benchmarks/)

## go-micro

- [Go Micro 中文文档 3.x](https://learnku.com/docs/go-micro/3.x)

- [Go Micro入门](http://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/GoMicro%E5%85%A5%E9%97%A8.html)

- [B站Golang后台架构](https://www.jianshu.com/p/b8e29b68252f?from=message&isappinstalled=0)
- [今日头条Go建千亿级微服务的实践](https://studygolang.com/articles/10117)
