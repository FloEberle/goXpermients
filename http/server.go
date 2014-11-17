package main

import (
    "net/http"
    "fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}


func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func (x *Struct) ServeHTTP (
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, x.Greeting + x.Punct + x.Who)
}


func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

    http.ListenAndServe("localhost:4000", nil)
}