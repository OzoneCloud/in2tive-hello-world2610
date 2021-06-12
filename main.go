package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `{"route":"/hello","response":"hello world v2 branch"}`)
}

func empty(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, `{"route":"/","branch":"v2"}`)
}

func main() {

	http.HandleFunc("/", empty)
	http.HandleFunc("/hello", hello)
	fmt.Print(http.ListenAndServe(":3000", nil))
}
