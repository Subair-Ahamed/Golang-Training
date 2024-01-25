package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleTCPConnection(conn net.Conn) {

	defer conn.Close()

	buffer, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("TCP Error reading:", err)
		return
	}

	fmt.Printf("TCP Received data from client: %v\n", buffer)

	msg := "Hello tcp!"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("TCP Error writing:", err)
		return
	}

	fmt.Println("data written", string(msg))
}

func main() {
	//listen on a TCP port
	tcpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening the TCP connection", err)
	}

	defer tcpListener.Close()

	fmt.Println("TCP Server listening on :8080")

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println("Error Accepting the TCP connection", err)
		}

		fmt.Printf("TCP Accepted connection from %s\n",conn.RemoteAddr())

		go handleTCPConnection(conn)
	}
}
