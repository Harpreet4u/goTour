package main

import (
    f "fmt"
    "path/filepath"
)

func main() {
    fileName := "/home/ubuntu/input.txt"
    f.Println(filepath.Base(fileName))  // input.txt

}
