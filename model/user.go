package model

import (
	"encoding/json"
	"html"
	"nutrition-api/database"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string      `gorm:"size:255;not null;unique" json:"username"`
	Password     string      `gorm:"size:255;not null;" json:"-"`
	Weight       json.Number `gorm:"size:255;" json:"weight"`
	Height       json.Number `gorm:"size:255;" json:"height"`
	Birhday_date *time.Time   `gorm:"size:255;" json:"birhday_date"`
	Sex          string      `gorm:"size:255;" json:"sex"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func DeleteAllUsers() error {
	err := database.Database.Where("ID > -1").Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func AllUsers () ([]User, error) {
	var users []User
	err := database.Database.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}