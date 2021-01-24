package doric

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Mux is a type that implements ServeHTTP
type Mux struct {
	router map[string]HandlerFunc
}

// New is the constructor of doric.mux
func New() *Mux {
	return &Mux{router: make(map[string]HandlerFunc)}
}

func (mux *Mux) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	mux.router[key] = handler
}

// Get defines the method to add GET request
func (mux *Mux) Get(pattern string, handler HandlerFunc) {
	mux.addRoute("GET", pattern, handler)
}

// Post defines the method to add POST request
func (mux *Mux) Post(pattern string, handler HandlerFunc) {
	mux.addRoute("POST", pattern, handler)
}

// Start defines the method to start a http server
func (mux *Mux) Start(addr string) (err error) {
	return http.ListenAndServe(addr, mux)
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := mux.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
