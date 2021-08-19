package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/decodeworms/citizenvisas/storage"
	"github.com/decodeworms/citizenvisas/types"
	"github.com/decodeworms/citizenvisas/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CitizenHandler struct {
	store *storage.DataStore
}

func NewCitizenHandler(s *storage.DataStore) CitizenHandler {
	return CitizenHandler{store: s}
}

func (c CitizenHandler) Create(w http.ResponseWriter, r *http.Request) {
	// setting the header for your endpoint is very important
	util.SetHeader(w)

	var data types.Citizen
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		e := errors.New(fmt.Sprintf("error decoding data: %s", err))
		json.NewEncoder(w).Encode(e)
	}

	log.Println("data being passed", data)

	err = c.store.Create(&data)
	if err != nil {
		// encode and return error. DO NOT log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(true)
}

func (c CitizenHandler) Citizen(w http.ResponseWriter, r *http.Request) {
	// setting the header for your endpoint is very important
	util.SetHeader(w)

	// extracting parameter from request
	params := mux.Vars(r)
	id := params["id"]

	log.Println("id being passed", id)

	// getting data from storage
	result, err := c.store.Citizen(id)
	if err != nil {
		// encode and return error. DO NOT log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(result)
}

func (c CitizenHandler) Citizens(w http.ResponseWriter, r *http.Request) {
	// setting the header for your endpoint is very important
	util.SetHeader(w)

	// getting data from storage
	citizens, err := c.store.Citizens()
	if err != nil {
		// encode and return error. DO NOT log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(citizens)
}

func (c CitizenHandler) Update(w http.ResponseWriter, r *http.Request) {
	// setting the header for your endpoint is very important
	util.SetHeader(w)

	var data types.Citizen
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		e := errors.New(fmt.Sprintf("error decoding data: %s", err))
		json.NewEncoder(w).Encode(e)
	}

	log.Println("data being passed", data)

	err = c.store.Update(data)
	if err != nil {
		// encode and return error. DO NOT log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(true)
}

func (c CitizenHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// setting the header for your endpoint is very important
	util.SetHeader(w)

	// extracting parameter from request
	params := mux.Vars(r)
	id := params["id"]

	log.Println("id being passed", id)

	err := c.store.Delete(id)
	if err != nil {
		// encode and return error. DO NOT log.Fatal(err)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(true)
}
