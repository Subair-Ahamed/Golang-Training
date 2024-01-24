//Dial - Establishes a network connection. It takes a network type ("tcp", "udp", etc.) and an address, and it returns a Conn interface and an error.

package main

import (
	"fmt"
	"net"
)

func main() {

	//Dial, DialTCP, DialUDP: functions to establish network connections
	conn, err := net.Dial("tcp", "sample.com:8080") //syntax: Dial(network, address string) (Conn, error)
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("connected to sample1.com")

}
