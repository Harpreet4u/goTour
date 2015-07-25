package main

import (
    "fmt"
    "crypto/sha1"
    "io"
    "log"
)

func main() {   
    hasher := sha1.New() // returns interface
    io.WriteString(hasher, "abc")
    b := []byte{}
    fmt.Printf("Result: %x\n", hasher.Sum(b))
    fmt.Printf("Result: %d\n", hasher.Sum(b))

    hasher.Reset()
    data := []byte("abc")
    n, err := hasher.Write(data)
    if n!=len(data) || err!= nil {
        log.Printf("Hash write error: %v / %v", n, err)
    }
    checksum := hasher.Sum(b)
    fmt.Printf("Result: %x\n", checksum)
}
    
