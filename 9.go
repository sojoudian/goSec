package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

var {
	iface = "en0"
	devFound = false
}

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panic(err)
	}
	for _, device := range devices {
		if dev.Name == iface {
			devFound = true
		}
		if !devFound {
			log.Panicf("Device not found %s", iface)
		}
	}
}
