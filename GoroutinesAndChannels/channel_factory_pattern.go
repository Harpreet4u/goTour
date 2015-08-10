package main
import (
    f "fmt"
    t "time"
)

func main() {
    stream := pump()
    go suck(stream)
    // or go suck(pump())
    t.Sleep(1e9)
}

func pump() chan int {
    ch := make(chan int)
    go func() {
        for i:=0;;i++ {
            ch <- i
        }
    }()
    return ch
}

func suck(ch chan int) {
    for {
        f.Println(<-ch)
    }
}
