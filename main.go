package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.elastic.co/apm/module/apmgorilla"
)

func hello(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	if rand.Float32() > 0.9 {
		fmt.Println("500 - Something bad happened!")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		fmt.Println("200 - Success")
		fmt.Fprintf(w, `{"route":"/hello","response":"hello world v3 branch"}`)
	}
}

func empty(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	if rand.Float32() > 0.9 {
		fmt.Println("500 - Something bad happened!")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		fmt.Println("200 - Success")
		fmt.Fprintf(w, `{"route":"/","branch":"v3"}`)
	}
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
