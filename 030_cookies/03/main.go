package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}
	counter, err := strconv.Atoi(c.Value)
	if err != nil {
		counter = 0
	}
	counter++
	c.Value = strconv.Itoa(counter)
	http.SetCookie(w, c)
	io.WriteString(w, c.Value)
}
