package main

import (
    "net/http"
)

func main() {

    // Create a new multiplexor object
    mux := http.NewServeMux()

    // Create a handler to serve HTTP requests with the contents of the stated file system
    // http.Dir("public") - I think this is a type conversion from the string public to
    // some type (Dir) that is recognized as the root of a file system
    fs := http.FileServer(http.Dir("public"))

    // Appears to be registering the path "/" with the handler fs
    // i.e. when "/" is requested (?), fs handles it (?)
    mux.Handle("/", fs)

    // This function starts listening on port 8080 and hooks up the multiplexor mux to
    // handle requests
    http.ListenAndServe(":8080", mux)

}
