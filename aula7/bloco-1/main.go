package main

import "net/http"

type Saudacao struct{ Msg string }

func (s Saudacao) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.Msg + "\n"))
}

func main() {
	http.ListenAndServe("127.0.0.1:8080", Saudacao{"oi"})
}
