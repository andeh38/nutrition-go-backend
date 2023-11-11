package model

import (
	"nutrition-api/database"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Type        string `gorm:"size:255;not null;default:simple" json:"type"`
	Name        string `gorm:"size:255;not null;" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
	ParentID    uint
	//Ingredients  *Food `json:"ingredients,omitempty"`
	Ingredients  *datatypes.JSONSlice[Food] `json:"ingredients,omitempty;default:null"`
	Weight       float64              `gorm:"size:255;default: 100 " json:"weight"`
	Calories     float64              `gorm:"size:255;not null " json:"calories"`
	Protein      float64              `gorm:"size:255;not null " json:"protein"`
	Fat          float64              `gorm:"size:255;not null " json:"fat"`
	Carbohydrate float64              `gorm:"size:255;not null " json:"carbohydrate"`
}

func (food *Food) Save() (*Food, error) {
	err := database.Database.Create(&food).Error
	if err != nil {
		return &Food{}, err
	}
	return food, nil
}

func FindFoodByName(name string, limit int) ([]Food, error) {
	var food []Food
	err := database.Database.Where("name LIKE ?", "%"+name+"%").Limit(limit).Find(&food).Error
	if err != nil {
		return []Food{}, err
	}
	return food, nil
}

func FindFoodById(id uint) (Food, error) {
	var food Food
	err := database.Database.Preload("Entries").Where("ID=?", id).Find(&food).Error
	if err != nil {
		return Food{}, err
	}
	return food, nil
}

func AllFood() ([]Food, error) {
	var foods []Food
	err := database.Database.Find(&foods).Error
	if err != nil {
		return []Food{}, err
	}
	return foods, nil
}
