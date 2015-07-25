package main
import (
    "io"
    "os"
    "fmt"
    "bufio"
    "flag"
)

func cat(r *bufio.Reader) {
    for {
        buf, err := r.ReadBytes('\n')
        if err == io.EOF {
            break
        }
        fmt.Fprintf(os.Stdout, "%s", buf)
    }
    return
}

func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        cat(bufio.NewReader(os.Stdin))
    }
    
    for i := 0; i < flag.NArg(); i++ {  
        f, err := os.Open(flag.Arg(i))
        if err != nil {
            continue
        }
        cat(bufio.NewReader(f))
    }
}
