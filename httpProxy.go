package main

import (
    "fmt"
    "net/http"
    "io"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintln(w, "hello world")
    vars := r.URL.Query();
    url := vars["url"][0]
    fmt.Println("url", url)
    resp, err := http.Get(url)
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    w.WriteHeader(resp.StatusCode)

    w.Write(body)
}

func main() {
    fmt.Println("started.")
    http.HandleFunc("/302", IndexHandler)
    err:=http.ListenAndServe("127.0.0.1:8888", nil)
    if err != nil {
        fmt.Println("err:", err)
    }
}