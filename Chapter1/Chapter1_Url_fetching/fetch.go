package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http://"){
            url = "http://" + url
        }
        resp, err := http.Get(url)
        fmt.Println("Code-status: ", resp.Status)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        dist := os.Stdout
        b, err := io.Copy(dist, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}
