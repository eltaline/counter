# Counter

Atomic Go Counters

# Usage

```package main

import (
        "fmt"
        "sync"

        "github.com/eltaline/counter"
)

type Dbs struct {
        Cnt *counter.Cint64
}

var (
        TestMap struct {
                Test map[string]*Dbs
        }
)

func main() {

        // Test with struct

        TestMap.Test = make(map[string]*Dbs)

        dbs, ok := TestMap.Test["one"]
        if !ok {
                TestMap.Test["one"] = new(Dbs)
                dbs = TestMap.Test["one"]
                dbs.Cnt = counter.NewInt64()
        }

        dbs.Cnt.Incr()
        dbs.Cnt.Decr()

        fmt.Println("Struct Thread Safe Value:", dbs.Cnt.Get())
        fmt.Println("Struct Thread Unsafe Value:", dbs.Cnt.Counter)

        // Direct use test

        c := counter.NewInt64()
        c.Incr()

        var wg sync.WaitGroup
        for i := 0; i < 8; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        for i := 0; i < 10000; i++ {
                                c.Incr()
                        }
                }()
        }
        wg.Wait()

        c.Decr()
        c.Decr()

        c.Add(int64(101))
        c.Sub(int64(100))

        fmt.Println("Direct Thread Safe Value", c.Get())
        fmt.Println("Direct Thread Unsafe Value", c.Counter)

}
```
