package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, you've requested: " + r.URL.Path + "\n"))
    })

    http.ListenAndServe(":8090", nil)
}