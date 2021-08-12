package main

import (
	"citizenvisas/handlers"
	"citizenvisas/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB
var txt string

// func init() {
// 	db, text = driver.ConnToDb()
// 	fmt.Print(text)
// 	host := os.Getenv("DB_HOST")

// }

func main() {

	router := mux.NewRouter()
	db, txt = storage.ConnToDb()
	fmt.Println(txt)

	citizenController := handlers.Citizen{}
	visaController := handlers.Visa{}

	router.HandleFunc("/create", citizenController.Create(db)).Methods("POST")
	router.HandleFunc("/citizen", citizenController.Citizen(db)).Methods("GET")
	router.HandleFunc("/citizens", citizenController.Citizens(db)).Methods("GET")
	router.HandleFunc("/update", citizenController.Update(db)).Methods("PUT")
	router.HandleFunc("/delete", citizenController.Delete(db)).Methods("DELETE")

	router.HandleFunc("/creates", visaController.Create(db)).Methods("POST")
	router.HandleFunc("/visas", visaController.Visas(db)).Methods("GET")
	router.HandleFunc("/visa", visaController.Visa(db)).Methods("GET")
	router.HandleFunc("/updateVisa", visaController.Update(db)).Methods("PUT")
	router.HandleFunc("/deleteVisa", visaController.Delete(db)).Methods("DELETE")
	// router.HandleFunc("/changeGender", citizenController.ChangeGender(db)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))

}
