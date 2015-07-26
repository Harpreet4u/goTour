package main

import (
    f "fmt"
    "path/filepath"
    "os"
)

func main() {   
    if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
        // returns Error object
        err := f.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
        f.Println(err)
        return
    }
}
