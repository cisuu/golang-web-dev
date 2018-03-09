package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		text := scanner.Text()
		io.WriteString(conn, text)
	}
	fmt.Println("code got here")
	io.WriteString(conn, "I see you connected\n")
	defer conn.Close()
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConnection(conn)
	}
}
