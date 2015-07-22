package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
        // Auto break here. Use fallthrough for no break behaviour
    case "linux":
        fmt.Println("Linux")
    default:
        // freebsd, openbsd
        // plan9, windows...
        fmt.Println("%s.", os)
    }
}
