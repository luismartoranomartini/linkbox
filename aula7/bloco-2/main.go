package main

import "net/http"

func ola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ola\n"))
}

func main() {
	http.ListenAndServe("127.0.0.1:8081", http.HandlerFunc(ola))
}
