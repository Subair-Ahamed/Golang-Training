package main

import (
	"fmt"
	"net"
)

//conn : a generic network connection interface
func handleConn(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	r, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading", err)
		return
	}

	data := buffer[:r]
	fmt.Println("Data read is ", string(data))

	message := "Hello server!"

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("error writing", err)
	}
	fmt.Println("Message written is ", message)
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening the connection", err)
		return
	}

	defer listener.Close()

	fmt.Println("server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting the connection", err)
			continue
		}
		fmt.Println("Accepted connection", conn.RemoteAddr()) //RemoteAddr is a method of the net.Conn interface in Go, and it returns the remote network address of the connection

		go handleConn(conn)
	}

}
