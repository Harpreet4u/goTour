package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(i)
        fmt.Println(s)
    }
}

func main() {
    // Runs until main block is not finished.
    // Asap main is finished goroutine will stop.
    go say("world")
    time.Sleep(1000*time.Millisecond)
    say("hello")
}
