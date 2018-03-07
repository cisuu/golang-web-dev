package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Loco-Key", "This is from Loco")
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(res, "<h1> Any code you want in this func </h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
