package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func roteador(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/saude":
		w.Write([]byte("ok\n"))
	case "/eco":
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		corpo, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(corpo))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           http.HandlerFunc(roteador),
		ReadHeaderTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
