package net

import (
	"fmt"
	"io"
	"net"
	"os"
)

func GetDial() {
	req := "GET / HTTP/1.1\n" + "Host: 127.0.0.1\n\n"
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	if _, err = conn.Write([]byte(req)); err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, conn)
	fmt.Println("OK")
}

func GetListen() {
	message := "Hello, I am a server"
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
