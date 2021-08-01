package main

import (
	"citizenvisas/driver"
	"citizenvisas/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB
var text string

func init() {
	db, text = driver.ConnToDb()
	fmt.Print(text)
	host := os.Getenv("DB_HOST")

}

func main() {

	router := mux.NewRouter()
	citizenController := handlers.Citizen{}

	router.HandleFunc("/create", citizenController.Create(db)).Methods("GET")
	// router.HandleFunc("/signup", citizenController.SignUp(db)).Methods("POST")
	router.HandleFunc("/confirm", citizenController.CheckIfAuserNameExistBeforeSignUp(db)).Methods("POST")
	router.HandleFunc("/citizen", citizenController.GetCitizen(db)).Methods("GET")
	router.HandleFunc("/citizens", citizenController.GetCitizens(db)).Methods("GET")
	router.HandleFunc("/search", citizenController.SearchUsingUserName(db)).Methods("GET")

	router.HandleFunc("/changeName", citizenController.ChangeName(db)).Methods("PUT")
	router.HandleFunc("/changeCountry", citizenController.ChangeCountry(db)).Methods("PUT")
	router.HandleFunc("/changeGender", citizenController.ChangeGender(db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))

}
