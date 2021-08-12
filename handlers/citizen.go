package handlers

import (
	"citizenvisas/storage"
	"citizenvisas/webservices"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Citizen struct {
}

var ctx context.Context

func (citizen Citizen) Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var fields webservices.Citizen
		json.NewDecoder(r.Body).Decode(&fields)

		var store = storage.DataStore{}
		err := store.Create(ctx, fields, db)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (citizen Citizen) Citizens(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fields []webservices.Citizen

		var store = storage.DataStore{}
		fields, err := store.Citizens(ctx, db)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(fields)
	}
}

func (citizen Citizen) Citizen(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var name webservices.Citizen
		json.NewDecoder(r.Body).Decode(&name)

		var store = storage.DataStore{}
		res, err := store.Citizen(ctx, db, name.Name)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(res)
	}
}

func (citizen Citizen) Update(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var data = storage.ChangeData{}

		var name webservices.Citizen
		json.NewDecoder(r.Body).Decode(&data)

		var store = storage.DataStore{}
		err := store.Update(ctx, db, name, data)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(name)
	}
}

func (citizen Citizen) Delete(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var name webservices.Citizen
		json.NewDecoder(r.Body).Decode(&name)

		var model webservices.Citizen

		var store = storage.DataStore{}
		err := store.Delete(ctx, db, name.Name, model)

		if err != nil {
			log.Fatal(err)
		}
	}
}
