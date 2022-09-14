package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"mysql-crud/config"
	"mysql-crud/model"
	"net/http"

	"github.com/gorilla/mux"
)

func AllBooks(w http.ResponseWriter, r *http.Request) {
	var book model.Books
	var response model.Response
	var booksArray []model.Books

	db := config.Connect()
	defer db.Close()
	config.Connect()

	rows, err := db.Query("SELECT * from Books")

	if err != nil {
		log.Print(err)
	}

	fmt.Println(rows)

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Isbn, &book.Firstname)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			booksArray = append(booksArray, book)
		}
	}

	fmt.Println(booksArray)

	response.Status = 200
	response.Message = "Success"
	response.Data = booksArray

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func FindBook(w http.ResponseWriter, r *http.Request) {
	var book model.Books
	var response model.Response
	var apiResponse []model.Books
	params := mux.Vars(r)

	db := config.Connect()
	defer db.Close()
	config.Connect()

	fmt.Println(params)

	rows, err := db.Query("SELECT * from Books where id = ?", params["id"])

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Isbn, &book.Firstname)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			apiResponse = append(apiResponse, book)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = apiResponse

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var newbook model.Books
	var response model.Response
	var apiResponse []model.Books

	db := config.Connect()
	defer db.Close()
	config.Connect()

	json.NewDecoder(r.Body).Decode(&newbook)
	fmt.Println(newbook)

	rows, err := db.Query("Insert into Books(id,title,isbn,firstname) values(?, ?, ?, ?)", &newbook.Id, &newbook.Title, &newbook.Isbn, &newbook.Firstname)
	if err != nil {
		log.Print(err)
	}

	fmt.Println(rows)

	for rows.Next() {
		err = rows.Scan(&newbook.Id, &newbook.Title, &newbook.Isbn, &newbook.Firstname)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			apiResponse = append(apiResponse, newbook)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = apiResponse

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var editbook model.Books
	var response model.Response
	var apiResponse []model.Books

	db := config.Connect()
	defer db.Close()
	config.Connect()

	json.NewDecoder(r.Body).Decode(&editbook)

	rows, err := db.Query("Update Books Set title = ? where id = ?", &editbook.Title, &editbook.Id)
	if err != nil {
		log.Print(err)
	}

	fmt.Println(rows, "dd")

	for rows.Next() {
		err = rows.Scan(&editbook.Id, &editbook.Title, &editbook.Isbn, &editbook.Firstname)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			apiResponse = append(apiResponse, editbook)
		}
	}

	fmt.Println(apiResponse)

	response.Status = 200
	response.Message = "Success"
	response.Data = apiResponse

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var deleteBook model.Books
	var response model.Response
	var apiResponse []model.Books

	db := config.Connect()
	defer db.Close()
	config.Connect()

	json.NewDecoder(r.Body).Decode(&deleteBook)

	_, err := db.Query("Delete from Books where id = ?", &deleteBook.Id)
	if err != nil {
		log.Print(err)
	}

	fmt.Println(apiResponse)

	response.Status = 200
	response.Message = "Success"
	response.Data = apiResponse

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
