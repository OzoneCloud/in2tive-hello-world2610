package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.elastic.co/apm/module/apmgorilla"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, `{"route":"/hello","response":"hello world v3 branch"}`)
}

func empty(w http.ResponseWriter, req *http.Request) {
	fmt.Println("empty")
	fmt.Fprintf(w, `{"route":"/","branch":"v3"}`)
}

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigFromEnvironment(),
	)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Use(apmgorilla.Middleware())

	r.HandleFunc(newrelic.WrapHandleFunc(app, "/", empty))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/hello", hello))
	fmt.Print(http.ListenAndServe(":3000", r))
}
