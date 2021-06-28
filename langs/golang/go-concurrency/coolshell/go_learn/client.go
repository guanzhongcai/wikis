package go_learn

import (
	"fmt"
	"net"
	"time"
)

func ClientRun()  {
	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := make([]byte, recv_buf_len)

	for i:=0; i<5; i++ {
		msg := fmt.Sprintf("Hello world, %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("writer error: %s", err.Error())
			break
		}
		fmt.Println("written OK:", msg)

		n, err = conn.Read(buf)
		if err != nil {
			println("read buffer error: ", err.Error())
			break
		}
		fmt.Println("read:", string(buf[0:n]))

		time.Sleep(time.Second)
	}
}
