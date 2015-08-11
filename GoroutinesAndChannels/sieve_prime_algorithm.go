// Pipes and filter pattern
package main
import f "fmt"

func generate(ch chan int) {
    for i := 2;;i++ {
        ch <- i
    }
}

func filter(in, out chan int, prime int) {
    for {
        i := <- in
        if i%prime != 0 {
            out <- i
        }
    }
}

func main() {
    ch := make(chan int)
    go generate(ch)
    //  [2..100]->[filter<2>]->[3,5,7,9,11,..]->[filter<3>]->[5,7,11]->...
    for {
        prime := <-ch
        f.Println(prime, " ")
        ch1 := make(chan int)
        go filter(ch, ch1, prime)
        ch = ch1
    }
}

