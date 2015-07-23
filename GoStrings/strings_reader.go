/*
    strings.NewReader(str string) *Reader
        NewReader returns a new Reader reading from s. It is similar to
        bytes.NewBufferString but more efficient and read-only.

        func (r *Reader) Read(b []byte) (n int, err error)
        func (r *Reader) ReadByte() (b byte, err error)
        func (r *Reader) ReadRune() (ch rune, size int, err error)

*/
package main

import (
    f "fmt"
    strs "strings"
    "io"
)

func main() {
    reader := strs.NewReader("Reader reads from here YO!")

    b := make([]byte, 8)
    for { 
        n, err := reader.Read(b)
        if err == io.EOF {
            break
        }    
        f.Printf("%s", b[:n])
    }
    
    reader = strs.NewReader("Reader reads from here YO!")
    
    for {
        b, err := reader.ReadByte()
        if err == io.EOF {
            break
        }
        f.Printf("%c", b)  // int read a time.
    }

    // Similarly ReadRune which reads int32 at a time.
}
