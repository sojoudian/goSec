package main

import (
	"fmt"
	"net"
)

func main() {
	IP := "scanme.nmap.org"
	//Port := "80"

	for i := 0; i < 100; i++ {
		//address := IP + ":" + string(i)
		address := fmt.Sprintf(IP+":%d", i)
		connection, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("[+] Connection established.. HOST and PORT: %v %v\n", i, connection.RemoteAddr().String())
		}

	}
}
