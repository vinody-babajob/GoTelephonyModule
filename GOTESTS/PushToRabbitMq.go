package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "./RabbitMqService"
)

func main() {
    http.HandleFunc("/", Index)

    http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello foo, %q", html.EscapeString(r.URL.Path))
        RabbitMqService.PublishMsg("FROM THE API")
    })


    log.Fatal(http.ListenAndServe(":8080", nil))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}