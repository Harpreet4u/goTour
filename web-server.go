/*
Package http serves HTTP requests using any value that implements http.Handler:

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
*/

package main

import (
    "fmt"
    "log"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello Go World!")
}

func main() {
    var h Hello
    err := http.ListenAndServe("localhost:8080", h)
    if err != nil {
        log.Fatal(err)
    }
}
