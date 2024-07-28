package main
import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    start := time.Now()
    fmt.Println(strings.Join(os.Args[1:], ""))
    fmt.Println(os.Args[1:])
    fmt.Println("executed time ", time.Since(start))
}
