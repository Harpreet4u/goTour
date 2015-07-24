/*
    func NewReader(rd io.Reader) *Reader
*/
package main

import (
    "bufio"
    f "fmt"
    "os"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)
    // ReadString(delimiter)
    f.Println("Enter your full name: ")

    // cannot use "\n"(string) as it needs '\n'(byte)
    input, err := inputReader.ReadString('\n')
    
    if err == nil {
        f.Println("The input was: ", input)
    }
}
    
