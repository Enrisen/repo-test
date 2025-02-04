// implement a feature that displays a simple message when the user visits the root URL of the web application.
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my first successfull go lang thingy!"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
