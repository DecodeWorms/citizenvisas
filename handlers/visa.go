package handlers
//
//import (
//	"encoding/json"
//	"github.com/decodeworms/citizenvisas/storage"
//	"github.com/decodeworms/citizenvisas/types"
//	"log"
//	"net/http"
//
//	"gorm.io/gorm"
//)
//
//type Visa struct{}
//
////var ctx context.Context
//
//func (visa Visa) Create(db *gorm.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var data types.Visa
//		json.NewDecoder(r.Body).Decode(&data)
//
//		var store = storage.VisaDataStore{}
//		store.Create(ctx, data, db)
//	}
//}
//
//func (visa Visa) Visas(db *gorm.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var fields []types.Visa
//
//		var store = storage.VisaDataStore{}
//		fields, err := store.Visas(ctx, db)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//		json.NewEncoder(w).Encode(fields)
//	}
//}
//
//func (visa Visa) Visa(db *gorm.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var field types.Visa
//		json.NewDecoder(r.Body).Decode(&field)
//
//		var store = storage.VisaDataStore{}
//		result, err := store.Visa(ctx, db, field.VisaNumber)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//		json.NewEncoder(w).Encode(result)
//	}
//}
//
//func (visa Visa) Update(db *gorm.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var data = storage.Data{}
//
//		var name types.Visa
//		json.NewDecoder(r.Body).Decode(&data)
//
//		var store = storage.VisaDataStore{}
//
//		err := store.Update(ctx, db, name, data)
//		if err != nil {
//			log.Fatal(err)
//		}
//		json.NewEncoder(w).Encode(name)
//
//	}
//}
//
//func (visa Visa) Delete(db *gorm.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var data types.Visa
//
//		var res types.Visa
//		json.NewDecoder(r.Body).Decode(&res)
//
//		var store = storage.VisaDataStore{}
//		err := store.Delete(ctx, db, data, res.VisaNumber)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}
