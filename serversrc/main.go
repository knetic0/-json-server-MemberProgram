package main

import (
	"fmt"
	"log"
	"os/exec"
)

const start string = "json-server"
const watch string = "--watch"
const file string = "../db.json"

func main() {
	fmt.Println("Server is starting...", "You can connect with http://localhost:3000/member")
	command := exec.Command(start, watch, file)
	fmt.Println(command.Output())
	stdout, err := command.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stdout)
}
