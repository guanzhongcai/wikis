package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	host string
	port string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 改写路由
	remote, err := url.Parse("http://" + this.host + ":" + this.port+"/api")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}
func startServer() {
	//被代理的服务器host和port
	h := &handle{host: "127.0.0.1", port: "8080"}
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
func main() {
	startServer()
}

/**
原始请求：
curl http://localhost:8888/hello
经过反向代理，发给后端服务请求：
curl http://localhost:8080/api/hello
 */