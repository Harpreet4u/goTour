package main

import "fmt"

func main() {
    defer fmt.Println("Runs in last and used in closing sockets. connections and files. Follows LIFO order execution.")
    
    fmt.Println("Hello, checkout defer functionality.")
}
    
    
