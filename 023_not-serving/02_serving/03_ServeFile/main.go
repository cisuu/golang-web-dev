package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/loco.png", locoPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="loco.png">
	`)
}

func locoPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "loco.jpg")
}
