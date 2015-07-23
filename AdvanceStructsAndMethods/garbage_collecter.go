package main

import (
    f "fmt"
    "runtime"
)

func main() {
    a := 8
    f.Println(a)
    m := &runtime.MemStats{}
    runtime.ReadMemStats(m)
    f.Printf("Memory used by go(Kbs): %d\n", m.Alloc/1024)
    runtime.GC()
    runtime.ReadMemStats(m)
    f.Printf("Memory used after GC by go(Kbs): %d\n", m.Alloc/1024)
}
