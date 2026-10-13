package main

import "net/http"

func criado(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Origem", "bloco-3")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("criado\n"))
}

func main() {
	http.ListenAndServe("127.0.0.1:8082", http.HandlerFunc(criado))
}
