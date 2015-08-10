type Empty interface {}

var empty Empty
...
data := make([]float64, N)
res := make([]float64, N)
sem := make([]float64, N)
...
for i, xi := range data {
    go func(i int, xi float64) {
        res[i] = doSomething(i, xi)
        sem <- empty
    }(i, xi)
}

// wait for goroutines to finish
for i:=0;i<N;i++ {<-sem}
