package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func ola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ola\n"))
}

func main() {

	server := http.Server{
		Addr:              "127.0.0.1:8085",
		Handler:           http.HandlerFunc(ola),
		ReadHeaderTimeout: 2 * time.Second,
	}

	err := server.ListenAndServe()
	fmt.Fprintln(os.Stderr, err)
}
