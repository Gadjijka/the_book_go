package main

import (
    "fmt"
    "os"
    "time"
)

func main(){
    start := time.Now()
    var s, sep string
    fmt.Println(os.Args[0])
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
        fmt.Println(i)
        fmt.Println(os.Args[i])
    }
    fmt.Println(s)
    fmt.Println("Затраченное время = ", time.Since(start))
}
