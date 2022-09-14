package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"mysql-crud/controller"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getBooks", controller.AllBooks).Methods("GET")
	router.HandleFunc("/getBook/{id}", controller.FindBook).Methods("GET")
	router.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	router.HandleFunc("/updateBook/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/deleteBook/{id}", controller.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
