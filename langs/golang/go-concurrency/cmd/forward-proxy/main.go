package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Pxy struct {
}
// swagger:operation PUT /api/v1/addr/update/{id} addr del
// ---
// summary: 修改用户地址
// description: 修改指定用户的addr
// parameters:
// - name: token
//   in: header
//   description: token
//   type: string
//   required: true
// - name: id
//   in: path
//   description: 地址id
//   type: string
//   required: true
// - name: addr
//   in: body
//   description: addr
//   type: string
//   required: true
// responses:
//   200: repoResp
//   400: badReq
func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	fmt.Printf("正向代理接收到请求：%s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport
	outReq := new(http.Request)
	// 浅拷贝
	*outReq = *req
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	//把消息利用该正向代理服务器发送给下游
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	//把下游得到的数据返回给上游
	//复制头部信息
	for key, value := range res.Header {
		for _, v := range value{
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)

	io.Copy(rw, res.Body)
	res.Body.Close()
}

type Server struct {

}

func (s* Server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("world"))
}

func main()  {
	http.Handle("/", &Pxy{})
	http.Handle("/api/hello", &Server{})
	http.ListenAndServe(":8080", nil)
}
