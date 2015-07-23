package main

import f "fmt"

type File struct{
    fd int
    name string
}

func main() {
    file := NewFile(12, "hello")
    f.Println(file)
}

func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    return &File{fd, name}
}
