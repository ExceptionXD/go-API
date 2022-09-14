package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"mysql-crud/controller"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getBooks", controller.AllBooks).Methods("GET")
	router.HandleFunc("/getBook/{id}", controller.FindBook).Methods("GET")
	router.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	router.HandleFunc("/updateBook/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/deleteBook/{id}", controller.DeleteBook).Methods("DELETE")

	fmt.Println("Connected to port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
