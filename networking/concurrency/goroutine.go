package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	
	defer conn.Close()

	fmt.Println("Accepted connection from", conn.RemoteAddr())
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle each connection concurrently using goroutine
		go handleConnection(conn)
	}
}