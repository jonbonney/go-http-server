package main

import (
    "net/http"
)

func main() {
    // Set up file server handler
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    // Start server
    println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        println(err.Error())
    }
}
