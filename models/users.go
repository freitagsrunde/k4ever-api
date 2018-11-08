package models

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func (u *User) Bind(r *http.Request) error {
	return nil
}

func (u *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUser(name string) User {
	return User{Name: name}
}

func ListUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return []User{}, err
	}
	return users, nil
}

func GetUser(id string, db *gorm.DB) (User, error) {
	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func CreateUser(user *User, db *gorm.DB) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
