package main
import f "fmt"

func main() {
    ch1 := make(chan int)
    go pump(ch1)
    // Blocked after printing 0
    f.Println(<-ch1)
}

func pump(ch chan int) {
    for i:=0;;i++ {
        ch <- i
    }
}
