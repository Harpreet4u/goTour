package main

import "fmt"

func main() {
    a, b := "Hello", "World"
    fmt.Println("Before swap:", a, b)
    a, b = b, a
    fmt.Println("After swap:", a, b)
}
