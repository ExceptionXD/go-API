package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sendHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sendHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
