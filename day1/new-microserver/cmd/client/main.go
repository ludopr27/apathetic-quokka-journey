package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    resp, err := http.Get("http://localhost:8080/hello")
    if err != nil {
        log.Fatalf("client error: %v", err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
