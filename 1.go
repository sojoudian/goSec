package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	ls := "ls"
	cmdObj := exec.Command(ls)
	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr
	err := cmdObj.Run()
	if err != nil {
		log.Fatal(err)
	}
}
