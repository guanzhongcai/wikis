package go_learn

import (
	"fmt"
	"io"
	"net"
)

const recv_buf_len = 1024

func ServerRun() {
	listener, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		panic("error listening: " + err.Error())
	}
	fmt.Println("starting the server")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error accept: " + err.Error())
		}
		fmt.Println("accepted the connection: ", conn.RemoteAddr())
		go EchoServer(conn)
	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, recv_buf_len)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Printf("warning: end of data: %s\n", err.Error())
			return
		default:
			fmt.Printf("error: reading data: %s\n", err.Error())
			return
		}
	}

}
