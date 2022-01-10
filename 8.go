package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panic(err)
	}
	for index, dev := range devices {
		fmt.Println(index, " ", dev.Name)
		for _, addr := range dev.Addresses {
			fmt.Println("\tIP: ", addr.IP)
			fmt.Println("\tINM: ", addr.Netmask)
		}
	}
}
