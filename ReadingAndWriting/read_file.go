/*
    Files in GO are pointers to objects of type os.File called filehandles.
    os.Stdin and os.Stdout are both of type *os.File
*/
package main

import (
    "bufio"
    f "fmt"
    "io"
    "os"
)

func main() {
    inputFile, inputErr := os.Open("input.txt")
    if inputErr != nil {
        f.Println("An error occurred on opening the file")
        return
    }
    defer inputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerErr := inputReader.ReadString('\n')
        if readerErr == io.EOF {
            return
        }
        f.Println(inputString)
    }
}

    
