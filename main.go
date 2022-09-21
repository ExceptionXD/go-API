package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"mysql-crud/controller"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/addUser", controller.AddUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/getUser/{email}", controller.GetUser).Methods("GET", "OPTIONS")

	c := cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
