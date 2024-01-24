package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolving IP address for a domain name
	ipAddr, err := net.ResolveIPAddr("ip", "example.com")
	if err != nil {
		fmt.Println("Error resolving IP address:", err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", "example.com:80")
	if err != nil {
		fmt.Println("Error resolving TCP address:", err)
		return
	}

	fmt.Printf("Resolved IP address: %s\n", ipAddr)

	fmt.Printf("Resolved TCP address: %s\n", tcpAddr)
}
