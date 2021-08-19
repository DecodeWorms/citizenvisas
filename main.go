package main

import (
	"context"
	"fmt"
	"github.com/decodeworms/citizenvisas/config"
	"github.com/decodeworms/citizenvisas/handlers"
	"github.com/decodeworms/citizenvisas/storage"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var appPort string
var dataStore storage.DataStore

func init() {
	// loads environment variables
	_ = godotenv.Load()

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	appPort = os.Getenv("APP_PORT")

	// configuration initialization
	cfg := config.Config{
		DatabaseName:     dbName,
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseUsername: dbUsername,
		DatabasePassword: dbPassword,
	}

	// context initialization
	initContext := context.Background()

	// db initialization
	client := storage.NewConnection(initContext, cfg)

	dataStore = storage.NewDataStore(initContext, client)
}

func main() {
	router := mux.NewRouter()

	citizenHandler := handlers.NewCitizenHandler(&dataStore)

	router.HandleFunc("/citizen/create", citizenHandler.Create).Methods("POST")
	router.HandleFunc("/citizen/{id}", citizenHandler.Citizen).Methods("GET")           // example: http://localhost:8080/citizen/1
	router.HandleFunc("/citizens", citizenHandler.Citizens).Methods("GET")
	router.HandleFunc("/citizen/update", citizenHandler.Update).Methods("PUT")
	router.HandleFunc("/citizen/delete/{id}", citizenHandler.Delete).Methods("DELETE")  // example: http://localhost:8080/citizen/delete/1

	//router.HandleFunc("/creates", visaController.Create(db)).Methods("POST")
	//router.HandleFunc("/visas", visaController.Visas(db)).Methods("GET")
	//router.HandleFunc("/visa", visaController.Visa(db)).Methods("GET")
	//router.HandleFunc("/updateVisa", visaController.Update(db)).Methods("PUT")
	//router.HandleFunc("/deleteVisa", visaController.Delete(db)).Methods("DELETE")
	// router.HandleFunc("/changeGender", citizenController.ChangeGender(db)).Methods("PUT")

	log.Println(fmt.Sprintf("starting server on port %s", appPort))
	log.Fatal(http.ListenAndServe(appPort, router))
}
