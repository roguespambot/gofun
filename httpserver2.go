// A slightly more complicated http server in golang

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Goodbye, world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/goodbye", goodbye)

	http.ListenAndServe(":8000", mux)
}
