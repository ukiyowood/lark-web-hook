package main

import (
	"fmt"
	"net"
	"os"
)

var port = "12321"

func main() {
	// 监听指定端口（例如：8080）
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error creating TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port " + port + "...")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 处理连接
		go handleConnection(conn)
	}
}

// 处理客户端连接
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {
		// 读取客户端发送的数据
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			break
		}

		// 打印接收到的数据
		fmt.Printf("Received: %s", string(buffer[:n]))

		// 将数据回显给客户端
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			break
		}
	}

	fmt.Println("Client disconnected:", conn.RemoteAddr())
}
