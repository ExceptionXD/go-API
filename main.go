package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func sendHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sendHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
