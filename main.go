package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `{"route":"/hello","response":"hello world main branch"}`)
}

func empty(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, `{"route":"/","branch":"main"}`)
}

func main() {

	http.HandleFunc("/", empty)
	http.HandleFunc("/hello", hello)
	fmt.Print(http.ListenAndServe(":3000", nil))
}