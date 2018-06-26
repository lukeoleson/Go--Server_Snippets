
package main

import (
    "fmt"
    "log"
    "net/http"
)

// A struct (object) to be used with our custom handler
type messageHandler struct {
    message string
}

// In order for our struct (messageHandler) to be an implementation of the
// interface http.Handler we need to override the interfaces ServeHTTP method
func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    // writes a string (m.message) to a writer (w)
    fmt.Fprintf(w, m.message)
}

func main() {

    // Create a multiplexor
    mux := http.NewServeMux()

    // Create an instance of our struct
    mh1 := &messageHandler{"Welcome to Go Web Development"}

    // Register a handler with the multiplexor
    // When the request is "/welcome", the handler mh1 is called
    // mh1 is a messageHandler object that implements the ServeHTTP method, which
    // means that really that method is now hooked up to this request
    mux.Handle("/welcome", mh1)

    // Again, just to show that the same handler can be used for different requests
    mh2 := &messageHandler{"net/http is awesome"}
    mux.Handle("/message", mh2)

    //
    log.Println("Listening...")

    // Listen on port 8080 and serve requests via the multiplexor mux
    http.ListenAndServe(":8080", mux)
}
