package main
import (
    f "fmt"
    t "time"
)

func main() {
    suck(pump())
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
    go func() {
        for v := range ch {
            f.Println(v)
        }
    }()
}
