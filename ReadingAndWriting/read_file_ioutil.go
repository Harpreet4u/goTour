package main

import (
    f "fmt"
    "io/ioutil"
    "os"
)

func main() {
    inputBuf, inputErr := ioutil.ReadFile("input.txt") // returns []byte
    if inputErr != nil {
        f.Fprintf(os.Stderr, "An error occurred on opening the file %s\n", inputErr)
        return
    }
    f.Printf("%s\n", string(inputBuf))
    err := ioutil.WriteFile("output.txt", inputBuf, 0644)
    if err != nil {
        panic(err.Error())
    }    
}

    
