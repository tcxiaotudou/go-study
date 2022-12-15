package main

import (
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Println("listen error: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			return
		}
		// 处理连接
		go Handle(conn)
	}

}

func Handle(conn net.Conn) {
	defer conn.Close()
	packet := make([]byte, 1024)
	for {
		// 阻塞直到读取数据
		n, err := conn.Read(packet)
		if err != nil {
			log.Println("read socket error: ", err)
			return
		}
		// 阻塞直到写入数据
		_, _ = conn.Write(packet[:n])
	}

}
