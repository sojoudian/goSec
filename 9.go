package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

var {
	iface = "en0"
	devFound = false
	snaplen = int32(1600)
	timeout = pcap.BlockForever
	filter = "tcp and port 443"
	promisc = false
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
		handle, err := pcap.OpenDev(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Panic(err)
		}
		// TODO:
		defer handle.Close()
		err := handle.SetBPFFilter(filter)
		if err != nil {
			log.Panic(err)
		}

		src := gopacket.NewPacketSource(handle, handle.LinkType())
		for pkt := range src.Packets {
			fmt.Println(pkt)
		}
	}
}
