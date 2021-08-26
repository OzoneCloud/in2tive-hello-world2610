package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `{"route":"/hello","response":"hello world v3 branch"}`)
}

func empty(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, `{"route":"/","branch":"v3"}`)
}

func main() {
	r := mux.NewRouter()
	r.Use(apmgorilla.Middleware())

	r.HandleFunc("/", empty)
	r.HandleFunc("/hello", hello)
	fmt.Print(http.ListenAndServe(":3000", r))
}
