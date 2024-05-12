package main

// http://localhost:8080/time

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"
)

type Time struct {
    CurrentTime string `json:"current_time"`
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
        currentTime := []Time{
            {CurrentTime: time.Now().Format(time.RFC3339)},
        }
        json.NewEncoder(w).Encode(currentTime)
    })

    http.ListenAndServe(":8080", nil)
}