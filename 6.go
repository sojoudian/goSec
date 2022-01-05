package main

import (
	"fmt"
	"net"
)

func scanPort(port int) {
	IP := "scanme.nmap.org"
	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)
	if err == nil {
		fmt.Printf("[+] Connection established.. HOST and PORT: %v %v\n", port, connection.RemoteAddr().String())
	}
}

func main() {
	for i := 0; i < 100; i++ {
		go scanPort(i)
	}
}
