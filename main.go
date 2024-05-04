package main

import (
    "net/http"
    "github.com/Wilovy09/Go_desde_0/utils"
)

func main() {
    http.HandleFunc("/", utils.Handler)
    http.ListenAndServe(":8080", nil)
}