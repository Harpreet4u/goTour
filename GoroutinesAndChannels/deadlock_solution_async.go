package main
import f "fmt"
import t "time"    

func f1(in chan int) {
    f.Println(<-in)
}

func main() {
    out := make(chan int, 2)
    // Does not block until buffer is filled.
    // Useful for scalabilty or sudden increase in requests
    out <- 2
    go f1(out)
    t.Sleep(1e9)
}
