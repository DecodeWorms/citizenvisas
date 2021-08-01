package handlers

import (
	"citizenvisas/repositories"
	"citizenvisas/types"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Citizen struct {
}

func (citizen Citizen) Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var table types.Citizen
		controller := repositories.Citizen{}
		controller.Create(db, table)

	}
}

// func (citizen Citizen) SignUp(db *gorm.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var userParameter types.Citizen
// 		json.NewDecoder(r.Body).Decode(&userParameter)

// 		var serverResponse types.Citizen

// 		controller := repositories.Citizen{}
// 		controller.SignUp(db, userParameter).Scan(&serverResponse)
// 		json.NewEncoder(w).Encode(serverResponse)
// 	}
// }

func (citizen Citizen) GetCitizen(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Citizen
		json.NewDecoder(r.Body).Decode(&userParameter)

		var serverResponse types.Citizen

		controller := repositories.Citizen{}
		controller.GetCitizen(db, userParameter).Scan(&serverResponse)
		json.NewEncoder(w).Encode(serverResponse)
	}
}

func (citizen Citizen) GetCitizens(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cit, err := GetCitizens(db)
		if err != nil {
			return nil, err
		}

		json.NewEncoder(w).Encode(citizen)
	}
}

func (citizen Citizen) ChangeCountry(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameter types.Citizen

		var country repositories.ChangeData
		json.NewDecoder(r.Body).Decode(&country)

		var serverResponse types.Citizen

		result, err := controller.ChangeCountry(db, userParameter, country)
		if err != nil {
			return nil, err
		}

		json.NewEncoder(w).Encode(serverResponse)

	}
}

func (citizen Citizen) ChangeName(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var model types.Citizen

		var result repositories.ChangeData

		json.NewDecoder(r.Body).Decode(&result)

		var serverResponse types.Citizen

		controller := repositories.Citizen{}
		controller.ChangeName(db, model, result).Scan(&serverResponse)

		json.NewEncoder(w).Encode(serverResponse) // work on this line..

	}
}

func (citizen Citizen) ChangeGender(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var gender types.Citizen

		var userParameter repositories.ChangeData
		json.NewDecoder(r.Body).Decode(&userParameter)

		var serverResponse types.Citizen

		controller := repositories.Citizen{}
		controller.ChangeGender(db, gender, userParameter).Scan(&serverResponse)
		json.NewEncoder(w).Encode(serverResponse)
	}
}

func (citizen Citizen) ChangeAge(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var age types.Citizen

		var userParameter repositories.ChangeData
		json.NewDecoder(r.Body).Decode(&userParameter)

		var serverResponse types.Citizen

		controller := repositories.Citizen{}

		controller.ChangeAge(db, age, userParameter).Scan(&serverResponse)
		json.NewEncoder(w).Encode(serverResponse)
	}
}

func (citizen Citizen) SearchUsingUserName(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userParameter repositories.ChangeData
		json.NewDecoder(r.Body).Decode(&userParameter)

		var database types.Citizen

		type serverResponse struct {
			Name    string
			Country string
			Gender  string
		}

		var result = []serverResponse{}

		controller := repositories.Citizen{}
		controller.SearchUsingUserName(db, database, userParameter).Scan(&result)
		json.NewEncoder(w).Encode(result)
	}
}

func (citizen Citizen) CheckIfAuserNameExistBeforeSignUp(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userParameters types.Citizen
		json.NewDecoder(r.Body).Decode(&userParameters)

		controller := repositories.Citizen{}
		controller.CheckIfAuserNameExistBeforeSignUp(db, userParameters)
	}
}

// func (citizen Citizen) DeleteCitizen(db *gorm.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var userId types.Citizen
// 		json.NewDecoder(r.Body).Decode(&userId)

// 		var serverResponse types.Citizen

// 		controller := repositories.Citizen{}
// 		controller.DeleteCitizen(db, userId).Scan(&serverResponse)
// 		json.NewEncoder(w).Encode(serverResponse)
// 	}
// }
