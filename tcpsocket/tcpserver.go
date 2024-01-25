package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port 8000")

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Message Received:", string(message))

		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n")) //writes the modified (uppercase) message back to the client using conn.Write
	}

}
