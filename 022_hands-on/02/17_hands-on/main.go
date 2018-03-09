package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleIndex(conn net.Conn) {
	body := "<h1>Holy Cow this is GET INDEX</h1>"
	writeToConn(conn, body)
}

func handleGETApply(conn net.Conn) {
	body := "<h1>Holy Cow this is GET APPLY</h1>"
	writeToConn(conn, body)
}

func handlePOSTApply(conn net.Conn) {
	body := "<h1>Holy Cow this is POST APPLY</h1>"
	writeToConn(conn, body)
}

func writeToConn(conn net.Conn, body string) {
	fmt.Println("code got here")
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func route(conn net.Conn, method, uri string) {
	switch {
	case method == "GET" && uri == "/":
		handleIndex(conn)
	case method == "GET" && uri == "/apply":
		handleGETApply(conn)
	case method == "POST" && uri == "/apply":
		handlePOSTApply(conn)
	}

}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	firstLine := true
	var method, uri string

	for scanner.Scan() {
		text := scanner.Text()

		if firstLine {
			reqHeader := strings.Fields(text)
			method, uri = reqHeader[0], reqHeader[1]
			fmt.Println("Method: ", method)
			fmt.Println("Uri: ", uri)
			fmt.Println("Protocol: ", reqHeader[2])
		}
		if text == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		firstLine = false
		route(conn, method, uri)
	}

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
