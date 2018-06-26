
/*
    DefaultServeMux is the default multiplexor in the http package

    So, rather than define and declare a multiplexor like we have been doing,
    we can just use shortcut functions in the http package to work with
    DefaultServeMux
*/
package main

import (
    "fmt"
    "log"
    "net/http"
)

// A function with the appropriate header to allow it to be converted to a http.HandlerFunc type
func messageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Go Web Development - DefaultServeMux")
}

func main() {

    // This shortcut bypasses creating a new multiplexor and registering the handler
    // to that multiplexor 
    http.HandleFunc("/welcome", messageHandler)

    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}
