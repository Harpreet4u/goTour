package main
import f "fmt"
import t "time"    

func f1(in chan int) {
    f.Println(<-in)
}

func main() {
    out := make(chan int)
    // Blocks until receiver is available so deadlock
    out <- 2
    go f1(out)
    t.Sleep(1e9)
}
