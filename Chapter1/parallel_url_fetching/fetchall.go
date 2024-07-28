package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
    "errors"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }
    file, err := os.OpenFile("result.txt", os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        if errors.Is(err, os.ErrNotExist){
            file, err = os.Create("result.txt")
            if err != nil {
                fmt.Println("Возникла ошибка", err)
                os.Exit(1)
            }
        }
    }
    defer file.Close()
    for range os.Args[1:] {
        file.WriteString(<-ch + "\n")
//        file.WriteString("%.2fs elapsed\n", time.Since(start).Seconds())
        file.Sync()
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
