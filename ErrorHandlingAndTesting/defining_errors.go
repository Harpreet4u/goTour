package main

import (
	err "errors"
	f "fmt"
)

var errNotFound error = err.New("Not found error")

func main() {

	f.Printf("Error: %v\n", errNotFound)
}
