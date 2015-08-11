package main
import f "fmt"
import t "time"

func main() {
    var c = make(chan int) // bidirectional
    go source(c)
    go sink(c)
    t.Sleep(1e9)
}

func source(ch chan<- int) {
    for { ch <- 1 }
}

func sink(ch <-chan int) {
    for { f.Println(<-ch) }
}


