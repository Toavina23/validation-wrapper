package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type User struct {
	Id       string    `json:"id"`
	Name     string    `validate:"required" json:"name"`
	Birthday time.Time `validate:"required" json:"birthday"`
}

var Validate *validator.Validate

func parseBody[U interface{}](w http.ResponseWriter, r *http.Request) (*U, error) {
	var data U
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		fmt.Fprint(w, "{message: bad request}")
		return nil, err
	}

	Validate = validator.New()
	err = Validate.Struct(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{message: bad request}")
		return nil, err
	}
	return &data, nil
}

func registerNewUser(w http.ResponseWriter, r *http.Request) {
	user, err := parseBody[User](w, r)
	if err != nil {
		return
	}
	fmt.Print(user)
}

func main() {
	fmt.Printf("test")
	r := mux.NewRouter()
	r.HandleFunc("/user", registerNewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

/*

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{message: internal server error}")
		return
	}
	user := &User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{message: bad request}")
		return
	}
	Validate = validator.New()
	err = Validate.Struct(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{message: bad request}")
		return
	}

*/
