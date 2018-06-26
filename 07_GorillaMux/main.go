/*
    Haven't finished - pg. 70
*/
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// The function with the appropriate header to be an implementation of http.HandlerFunc
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Get Note Handler")
}

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Post Note Handler")
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Put Note Handler")
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Delete Note Handler")
}


func main() {

    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
    r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
    r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
    r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

    log.Println("Listening...")

    server := &http.Server{
        Addr: ":8080",
        Handler: r,
    }
    server.ListenAndServe()
}
