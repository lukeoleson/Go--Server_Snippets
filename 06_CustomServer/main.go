/*
   This example shows us how to set up our own Server object to listen and serve
   rather than using the default server that we have been using in previous examples.
   This way we can set the properties of our Server object to custom settings
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// The function with the appropriate header to be an implementation of http.HandlerFunc
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development - Custom server settings")
}

func main() {

	// Here we are using the shortcut again (from ex. 05) from the http package
	// to just use DefaultServerMux and plumb the messageHandler HandlerFunc
	// to the /welcome request
	http.HandleFunc("/index.html", messageHandler)

	// This time we create a server object (struct) and make some custom settings
	// to its fields
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Our log
	log.Println("Listening...")

	// So now we need to call ListenAndServe on our custom server rather than
	// on the default server which we were using previously
	server.ListenAndServe()
}
