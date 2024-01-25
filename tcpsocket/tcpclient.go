package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin) //bufio.NewReader reads from the standard input (os.Stdin)
		fmt.Print("Text to send: ")

		text, _ := reader.ReadString('\n') 
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n') //it reads the response from the server as a string, until a newline ('\n') character is encountered
		fmt.Print("Message from server: " + message)
	}

}