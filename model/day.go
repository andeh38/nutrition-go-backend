package model

import (
	"nutrition-api/database"

	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Day struct {
	gorm.Model
	Date      *time.Time `gorm:"size:255;" json:"date"`
	UserID    uint
	Breakfast *datatypes.JSONSlice[Food] `json:"breakfast,omitempty;default:null"`
	Brunch    *datatypes.JSONSlice[Food] `json:"brunch,omitempty;default:null"`
	Lunch     *datatypes.JSONSlice[Food] `json:"lunch,omitempty;default:null"`
	Dinner    *datatypes.JSONSlice[Food] `json:"dinner,omitempty;default:null"`
}

func (day *Day) Save() (*Day, error) {
	err := database.Database.Create(&day).Error
	if err != nil {
		return &Day{}, err
	}
	return day, nil
}

func AllDays() ([]Day, error) {
	days := []Day{}
	err := database.Database.Find(&days).Error
	if err != nil {
		return []Day{}, err
	}
	return days, nil
}

func SliceDays(minDate, maxDate *time.Time, userId uint) ([]Day, error) {
	days := []Day{}
	err := database.Database.Where("user_id = ?", userId).Where("date BETWEEN ? AND ?", minDate, maxDate).Order("date").Find(&days).Error
	if err != nil {
		return []Day{}, err
	}
	return days, nil
}

func (day *Day) FindOrCreate() (*Day, error) {
	y, m, d := day.Date.Date()
	*day.Date = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	err := database.Database.Where(Day{UserID: day.UserID, Date: day.Date}).FirstOrCreate(&day).Error
	if err != nil {
		return &Day{}, err
	}
	return day, nil
}

func (day *Day) Update() (*Day, error) {
	err := database.Database.Updates(&day).Error
	if err != nil {
		return &Day{}, err
	}
	return day, nil
}

func DeleteAll() error {
	err := database.Database.Where("ID > -1").Delete(&Day{}).Error
	if err != nil {
		return err
	}
	return nil
}
