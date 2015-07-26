package main

import f "fmt"

func main() {
    f.Println("Starting the program")
    panic("A severe error occurred: stopping the program!")
    f.Println("Ending the program")
}
