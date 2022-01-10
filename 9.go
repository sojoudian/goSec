package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface    = "en0"
	devFound = false
	snaplen  = int32(1600)
	timeout  = pcap.BlockForever
	filter   = "tcp and port 80"
	promisc  = false
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panic(err)
	}
	for _, dev := range devices {
		if dev.Name == iface {
			devFound = true
		}
		if !devFound {
			log.Panicf("Device not found %s", iface)
		}
		handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Panic(err)
		}
		// TODO:
		defer handle.Close()
		errNew := handle.SetBPFFilter(filter)
		if err != nil {
			log.Panic(errNew)
		}

		src := gopacket.NewPacketSource(handle, handle.LinkType())
		for pkt := range src.Packets() {
			//fmt.Println(pkt)
			applayer := pkt.ApplicationLayer()
			if applayer == nil {
				continue
			}
			payload := applayer.Payload()
			search_arr := []string{"name", "username", "password", "pass"}
			for _, s := range search_arr {
				index := strings.Index(string(payload), s)
				if index != -1 {
					fmt.Println(string(payload[index:index+100]), s)
				}
			}

		}
	}
}
