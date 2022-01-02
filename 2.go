package main

import (
	"log"
	"os"
	"os/exec"
)

func excCmd(cmd string, argsArr []string) (err error) {
	args := argsArr
	cmdObj := exec.Command(cmd, args...)
	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr
	// because error was initialized in the function definition ==>(err error)
	// func excCmd(cmd string, argsArr []string) (err error)
	//we do not need to use := onlu = is enough
	err = cmdObj.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	return nil

}

func main() {
	command := "ls"
	excCmd(command, []string{"-lah"})

}
