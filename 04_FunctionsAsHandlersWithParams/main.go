/*
    Same theory as 03_FunctionsAsHandlers, only this time we wrap our convertible function
    inside another function that can take in parameters
*/

package main

import (
    "fmt"
    "log"
    "net/http"
)

/*
    A function that takes in params and returns a function that can be converted to
    a http.HandlerFunc type (i.e. has the appropriate signature)

    This is using "closure" ... how/what?

*/
func messageHandler(message string) http.HandlerFunc {

    // Method #1 - as per the book
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, message)
    })


}

// Method #2 - Same thing basically but I just return a function and convert it to
// a HandlerFunc after I call this function.  I was hoping it may be a bit more
// obvious what's happening here, but it also is more complicated :)
func messageHandler2(message string) func(w http.ResponseWriter, r *http.Request) {

    // Inner functions do not get names, they get assigned to variables
    // of type function (I might have made that up)
    innerFunc := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, message)
    }

    return innerFunc
}


func main() {

    // Create a ServeMux object (multiplexor)
    mux := http.NewServeMux()

    // Get ahold of a
    mh := messageHandler("Welcome to Go Web Development - Handler Logic with Closure")

    mh2 := http.HandlerFunc(messageHandler2("Welcome to Go Web Development - Alt. Handler Logic with Closure"))

    // Register the handler with the multiplexor for a certain request
    mux.Handle("/welcome", mh)
    mux.Handle("/other", mh2)

    // Log the actions...
    log.Println("Listening...")

    // Listen to port 8080 and serve requests with the multiplexor
    http.ListenAndServe(":8080", mux)
}
