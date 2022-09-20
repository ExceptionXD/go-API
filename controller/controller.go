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

func AddUser(w http.ResponseWriter, r *http.Request) {
	var newbook model.User
	var response model.Response
	var apiResponse []model.User

	db := config.Connect()
	defer db.Close()
	config.Connect()

	json.NewDecoder(r.Body).Decode(&newbook)
	fmt.Println(newbook)

	rows, err := db.Query("Insert into Users(name, email , password,id) values(?, ?, ?, ?)", &newbook.Name, &newbook.Email, &newbook.Password, &newbook.Id)
	if err != nil {
		log.Print(err)
	}

	fmt.Println(rows)

	for rows.Next() {
		err = rows.Scan(&newbook.Name, &newbook.Email, &newbook.Password, &newbook.Id)
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var response model.Response
	var apiResponse []model.User
	params := mux.Vars(r)

	db := config.Connect()
	defer db.Close()
	config.Connect()

	fmt.Println(params)

	rows, err := db.Query("SELECT * from Users where email = ?", params["email"])

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&user.Name, &user.Email, &user.Password, &user.Id)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			apiResponse = append(apiResponse, user)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = apiResponse

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
