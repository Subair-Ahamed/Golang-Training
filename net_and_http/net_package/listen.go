//Listens for incoming connections on the local network address. It returns a Listener interface and an error. establishes a listener to the connection.

package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000") //syntax: Listen(network, address string) (Listener, error)

	if err != nil {
		fmt.Println("Listen error:", err)
	}

	defer listener.Close()

	fmt.Println("Server listening on port:8000")
}
