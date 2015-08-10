func compute(ch chan int) {
    ch < someComputation()
}

func main() {
    ch := make(chan int)
    go compute()
    doSomethingElseForAWhile()
    result := <-ch
}
