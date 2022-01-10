package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func main() {
	targetIP := "10.0.0.0/24"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(targetIP),
		nmap.WithPorts("80,443,8080"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatal("error: ", err)
	}

	results, warning, err := scanner.Run()
	if err != nil {
		log.Fatal("error: ", err)
	}
	if warning != nil {
		log.Fatal("warning: ", warning)
	}
	for _, host := range results.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		fmt.Printf("IP: %q\n", host.Addresses[0])
		if len(host.Addresses) > 1 {
			fmt.Printf("MAC %v\n", host.Addresses[1])
		}
		for _, port := range host.Ports {
			fmt.Printf("\t Port %d %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
}
