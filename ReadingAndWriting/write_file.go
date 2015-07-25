package main
import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    outputFile, err := os.OpenFile("output.dat", os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        fmt.Printf("An error occurred.")
        return 
    }
    defer outputFile.Close()

    outputWriter := bufio.NewWriter(outputFile)
    outputWriter.WriteString("Hello World!")
    outputWriter.Flush()

    // or
    // outputFile.WriteString("blah blah!")
}
