/*
Instead of implementing the http.Handler interface, we use the http.HandlerFunc type
to serve as a hanlder.

The http.HandlerFunc is a standard library type that has a built-in ServeHTTP method, and
therefore it satifies the http.Hander interface and can serve as a handler

Any function can be converted to a http.HandlerFunc function if the signature matches
--> func(http.ResponseWriter, *http.Request)
*/

package main

import (
    "fmt"
    "log"
    "net/http"
)

// A function with the appropriate header to allow it to be converted to a http.HandlerFunc type
func messageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Go Web Development - Functions as Handlers")
}

func main() {

    // Create a multiplexor
    mux := http.NewServeMux()

    // Create a handler
    // In this case we need to convert our regular function (defined above) to an
    // http.HandlerFunc
    // This seems crazy, but in Go functions are a type (I geuss) and they can be assigned
    // to aliases. So mh is actually referring to a function of type http.HandlerFunc now
    mh := http.HandlerFunc(messageHandler)

    // Register the handlers with the multiplexor
    mux.Handle("/welcome", mh)

    // // There is also a shortcut (probably preferred) that converts the messageHandler into
    // // a HandlerFunc automatically (just skips that last step) -->
    // mux.HandleFunc("/welcome", mh)


    // Log whats happening to standard out
    log.Println("Listening...")

    // Listen to port 8080 and serve requests through mux
    http.ListenAndServe(":8080", mux)
}
