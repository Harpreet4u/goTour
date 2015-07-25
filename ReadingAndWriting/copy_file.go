package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    CopyFile("target.txt", "input.txt")
    fmt.Println("Copy done!")
}

func CopyFile(d, s string) (written int64, err error) {
    src, err := os.Open(s)
    if err !=nil {
        return
    }
    defer src.Close()

    dst, err := os.OpenFile(d, os.O_WRONLY | os.O_CREATE, 0644)
    if err !=nil {
        return
    }
    defer dst.Close()
    
    return io.Copy(dst, src)
}
