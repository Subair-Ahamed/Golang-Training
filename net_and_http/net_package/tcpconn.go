package main

import (
	"fmt"
	"net"
)

func handleTCPConnection(conn *net.TCPConn) { //similar method for net.UDPConn

	defer conn.Close()

	// Reading from the TCP connection
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer) //for UDPConn : n,addr,err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("TCP Error reading:", err)
		return
	}

	received := buffer[:n]
	fmt.Printf("TCP Received data from client: %s\n", received) //for udp("%s,%s",addr,received)

	// Writing to the TCP connection
	message := "Hello, TCP client!"
	_, err = conn.Write([]byte(message)) //for udp : _,err = conn.WriteToUDP([]byte(message))
	if err != nil {
		fmt.Println("TCP Error writing:", err)
		return
	}

	fmt.Printf("TCP Sent message to client: %s\n", string(message))
}

func main() {
	// Listen on a TCP port
	//The net.ResolveIPAddr() function is used to resolve an IP address from a given network and address string.
	tcpAddr,_ := net.ResolveTCPAddr("tcp", ":8080")

	//listen on a TCP port
	tcpListener,_ := net.ListenTCP("tcp", tcpAddr) //for udp : net.ListenUDP()

	defer tcpListener.Close()

	fmt.Println("TCP Server listening on :8080")

	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("Error Accepting the TCP connection", err)
		}
		fmt.Printf("TCP Accepted connection from %s\n", conn.RemoteAddr())

		// Handle the TCP connection in a separate goroutine
		go handleTCPConnection(conn)
	}
}
