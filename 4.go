package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

//mac address changer

func excCmd(cmd string, argsArr []string) (err error) {
	args := argsArr
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr
	cmdObj.Stdin = os.Stdin

	err = cmdObj.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	return nil

}
func main() {
	iface := flag.String("iface", "", "Interface you want to change the physical address")
	newMac := flag.String("newMac", "", "Provide the new physical address")

	flag.Parse()

	// shutdown the wireless network
	//excCmd("sudo", []string{"ifconfig", "en0", "down"})
	//change the physical address
	//excCmd("sudo", []string{"ifconfig", "en0", "hw", "ether", "00:11:22:33:44:55"})
	excCmd("sudo", []string{"ifconfig", *iface, "ether", *newMac})

	//bring back the wireless network up
	excCmd("sudo", []string{"ifconfig", *iface, "up"})
}
