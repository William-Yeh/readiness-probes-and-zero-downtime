package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
    "time"
)

func main() {
    if len(os.Args) > 1 {
        if delay, err := strconv.Atoi(os.Args[1]); err == nil {
            fmt.Printf("Taking %d seconds to initialize... ", delay)
            time.Sleep(time.Duration(delay) * time.Second)
            fmt.Println("Done.")
        }
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "HELLO WORLD!")
    })

    http.HandleFunc("/health", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "OK")
    })

    http.ListenAndServe(":80", nil)
}