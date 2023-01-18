package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type persons struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createNewPerson(w http.ResponseWriter, r *http.Request) { //router handler
	reqBody, _ := ioutil.ReadAll(r.Body)
	var person persons
	json.Unmarshal(reqBody, &person)
	db.Create(&person)
	fmt.Println("Create New Person")
	json.NewEncoder(w).Encode(person)
}

func main() {
	db, err = gorm.Open("mysql", "root:Tungusonali@1102@tcp(127.0.0.1:3306)/sk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open!")
	} else {
		fmt.Println("Connection Established!")
	}
	defer db.Close()
	db.AutoMigrate(&persons{})
	router := mux.NewRouter()
	router.HandleFunc("/person", createNewPerson).Methods("POST")
	http.ListenAndServe(":8000", router)

}
