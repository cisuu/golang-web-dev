package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	firstLine := true

	for scanner.Scan() {
		text := scanner.Text()
		if firstLine {
			reqHeader := strings.Fields(text)
			fmt.Println("Method: ", reqHeader[0])
			fmt.Println("Uri: ", reqHeader[1])
			fmt.Println("Protocol: ", reqHeader[2])
		}
		if text == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		firstLine = false
	}
	body := "I see you connected\n"
	fmt.Println("code got here")
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

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

		go serve(conn)
	}
}
