for i, v := range data {
    go func(i int, v float64) {
        doSomething(i, v)
    }(i, v)
}

// huge performance gain 
