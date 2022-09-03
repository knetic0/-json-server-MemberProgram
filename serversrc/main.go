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
	channel := make(chan string)
	fmt.Println("Server is starting...", "You can connect with http://localhost:3000/member")
	command := exec.Command(start, watch, file)
	stdout, err := command.Output()
	go Addr(channel, string(stdout))
	ip_addr := <-channel
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(ip_addr)
	}
}

func Addr(channel chan string, stdout string) {
	channel <- stdout
}
