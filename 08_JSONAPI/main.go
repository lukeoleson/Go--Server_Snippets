/*
    I wouldn't say that this works exactly. I made some changes to make it work a tiny little bit.
    see my changes under "NOTE:" in the code...

    The only things that seem to work are POST (with some changes) and GET

    I don't know what the ID's are or how I would enter them
*/
package main

import (
    // "fmt"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "time"

    "github.com/gorilla/mux"
)

/*
    Data model and data store
*/

type Note struct {

    // I think this extra syntax with the 'json:...' is a tag (?) saying that these
    // attributes of this struct are in JSON ??? or will be encoded in JSON
    Title string `json:"title"`
    Description string `json:"description"`
    CreatedOn time.Time `json:"createdon"`
}

/*
    There is no database in this program for persistent storage, so we are using
    a map of (string, Note) pairs instead to store our notes (temporarily)
*/
// Store for the Notes collection
var noteStore = make(map[string]Note)

// Variable to generate keys for the collection
var id int = 0

// HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {

    var note Note

    // NOTE: I commented out this block as it was throwing an error
    // I don't understande exactly why we are decoding a note that we just created
    // here???
    /*
    // Decode the incoming Note json
    // err := json.NewDecoder(r.Body).Decode(&note)
    // if err != nil {
    //     fmt.Println("\n\nPostNoteHandler error\n\n")
    //     panic(err)
    // }
    */

    // NOTE: I added these two lines so that my entries would have something
    // in the title and description fields...
    note.Title = "This is a title"
    note.Description = "It is cool...apparently"

    note.CreatedOn = time.Now()
    id++
    k := strconv.Itoa(id)
    noteStore[k] = note

    j, err := json.Marshal(note)
    if err!= nil {
        panic(err)
    }
    w.Header().Set("Content - Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(j)
}

// HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
    var notes []Note
    for _, v := range noteStore {
        notes = append(notes, v)
    }
    w.Header().Set("Content - type", "application/json")
    j, err := json.Marshal(notes)
    if err != nil {
        panic(err)
    }
    w.WriteHeader(http.StatusOK)
    w.Write(j)
}

// HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
    var err error
    vars := mux.Vars(r)
    k:= vars["id"]
    var noteToUpd Note
    // Decode the incoming Note json
    err = json.NewDecoder(r.Body).Decode(&noteToUpd)
    if err != nil {
        panic(err)
    }
    if note, ok := noteStore[k]; ok {
        noteToUpd.CreatedOn = note.CreatedOn
        // delete existing item and add the updated item
        delete(noteStore, k)
        noteStore[k] = noteToUpd
    } else {
        log.Printf("Could not find key of Note %s to delete", k)
    }
    w.WriteHeader(http.StatusNoContent)
}

// HTTP Delete - api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    k := vars["id"]
    // remove from Store
    if _, ok := noteStore[k]; ok {
        // delete existing item
        delete(noteStore, k)
    } else {
        log.Printf("Could not find key of Note %s to delete", k)
    }
    w.WriteHeader(http.StatusNoContent)
}

// Entry point of the program
func main() {
    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
    r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
    r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
    r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

    server := &http.Server{
        Addr: ":8080",
        Handler: r,
    }

    log.Println("Listening...")
    server.ListenAndServe()
}
