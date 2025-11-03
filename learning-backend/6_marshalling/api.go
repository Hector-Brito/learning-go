package main

import (
	"encoding/json"
	"net/http"
	"errors"
)

type application struct {
	addr string
}

var users = []User{}


func (app *application) getUserHandler(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type","application/json")

	//encode users slice to json
	err := json.NewEncoder(writer).Encode(users)

	if err != nil{
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)

}

func (app *application) createUserHandler(writer http.ResponseWriter, request *http.Request){
	var payload User
	err := json.NewDecoder(request.Body).Decode(&payload)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		ID: payload.ID,
		Username: payload.Username,
	}

	if err := insertUser(u); err != nil{
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return 

	}
	writer.WriteHeader(http.StatusCreated)
}


func insertUser(u User) error {
	//input validation
	if u.Username == "" {
		return errors.New("Username is required")
	}
	if  u.ID == 0 {
		return errors.New("ID is required")
	}

	//storage validation
	for _, user := range users {
		if user.ID == u.ID || user.Username == u.Username{
			return errors.New("User already exists")
		}
	}
	users = append(users, u)
	return nil
}