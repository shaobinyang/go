package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func hello2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "myserver:"+r.URL.String())
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	s := http.Server{
		Addr:    ":9000",
		Handler: &myHander{},
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello
	mux["/h"] = hello2
	s.ListenAndServe()
}

type myHander struct{}

func (myHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "myserver:"+r.URL.String())
}
