package main

import (
	"fmt"
	"net/http"
)

func info(w http.ResponseWriter, r *http.Request) {
	cor := r.URL.Query().Get("cor")
	fmt.Fprintf(w, "%s %s cor=%q x=%q\n", r.Method, r.URL.Path, cor, r.Header.Get("X-Teste"))
}

func main() {
	http.ListenAndServe("127.0.0.1:8083", http.HandlerFunc(info))
}
