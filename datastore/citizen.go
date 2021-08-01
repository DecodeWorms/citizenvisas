package repositories

import (
	"citizenvisas/types"
	"fmt"

	"gorm.io/gorm"
)

type Citizen struct {
	client *gorm.DB
}

type ChangeData struct {
	Name       string
	NewName    string
	Country    string
	NewCountry string
	Age        int
	NewAge     int
	Gender     string
	NewGender  string
}

func NewCitizen()
func (citizen Citizen) Create(db *gorm.DB, model types.Citizen) error {

	return db.AutoMigrate(&model)

}

// func (citizen Citizen) SignUp(db *gorm.DB, model types.Citizen) *gorm.DB {
// 	return db.Create(&model)

// }

func (citizen Citizen) GetCitizen(db *gorm.DB, citizenId string) (*types.Citizen, error) {
	var m types.Citizen
	return &m, db.Where("id = ?", citizenId).Find(&m).Error
}

func (citizen Citizen) GetCitizens(db *gorm.DB) ([]*types.Citizen, error) {
	var m []*types.Citizen
	return m, db.Find(&m).Error
}

func (citizen Citizen) ChangeCountry(db *gorm.DB, model types.Citizen, result ChangeData) *gorm.DB {
	return db.Model(&types.Citizen{}).Where("country = ?", result.Country).Update("country", result.NewCountry)
}

func (citizen Citizen) ChangeName(db *gorm.DB, model types.Citizen, result ChangeData) *gorm.DB {

	return db.Model(&types.Citizen{}).Where("name = ?", result.Name).Update("name", result.NewName)
}

func (citizen Citizen) ChangeGender(db *gorm.DB, model types.Citizen, result ChangeData) *gorm.DB {
	return db.Model(&types.Citizen{}).Where("gender = ?", result.Gender).Update("gender", result.NewGender)
}

func (citizen Citizen) ChangeAge(db *gorm.DB, model types.Citizen, result ChangeData) *gorm.DB {

	return db.Model(&types.Citizen{}).Where("age = ?", result.Age).Update("age", result.NewAge)
}

func (citizen Citizen) SearchUsingUserName(db *gorm.DB, model types.Citizen, result ChangeData) *gorm.DB {

	return db.Model(&types.Citizen{}).Select("name,country,gender").Where("name LIKE ?", result.Name+"%").Find(&model)

}

func (citizen Citizen) CheckIfAuserNameExistBeforeSignUp(db *gorm.DB, model types.Citizen) *gorm.DB {

	var result types.Citizen
	db.Model(&types.Citizen{}).Select("name").Where("name = ?", model.Name).Find(&model).Scan(&result)

	if result.Name == model.Name {
		fmt.Println("name already exist")
		return db
	}
	return db.Create(&model)

}

// func (citizen Citizen) DeleteCitizen(db *gorm.DB, model types.Citizen) *gorm.DB {
// 	return db.Where("id = ?", model.ID).Delete(&model)
// }
