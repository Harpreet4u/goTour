package main

import (
    f "fmt"
    "os"
)

var user = os.Getenv("USER")

func check() {
    if user == "" {
        panic("Unknown user: no value for $USER")
    }
    f.Println(user)
}

func main() {
    check()
}
