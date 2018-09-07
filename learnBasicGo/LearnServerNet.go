package main

import (
	"fmt"
	"net"
)

/*net*/
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("read", string(buf[:n]))
	}

}

func main() {

	// 服务端
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
	}
	for {
		con, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go process(con)
	}

}
