/*
The io.Reader interface has a Read method:

func (T) Read(b []byte) (n int, err error)
Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.
*/

package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, Reader!!!!!")
    
    b := make([]byte, 8) // reading 8 bytes at a time.
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err == io.EOF {
            break
        }
    }
}
