package main

import (
    f "fmt"
    t "time"
)

func main() {
    ch := make(chan string)
    go sendData(ch)
    go getData(ch)
    t.Sleep(1e9)
}

func sendData(ch chan string) {
    ch <- "Mountain View"
    ch <- "Delhi"
    ch <- "Chandigarh"
    ch <- "Bangalore"
    ch <- "Mohali"
}

func getData(ch chan string) {
    var input string
    for {
        input = <- ch
        f.Printf("%s \n", input)
    }
}
