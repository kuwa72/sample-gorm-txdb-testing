package usecase

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint `gorm:"primaryKey"`
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
)

func LoginUser(db *gorm.DB, email, password string) (*User, error) {
	var user User
	err := db.First(&user, "email = ? and password = ?", email, password).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &user, nil
}

func CreateUser(db *gorm.DB, name string, email string, password string) (*User, error) {
	var user User
	err := db.First(&user, "email = ?", email).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		break
	case err != nil:
		return nil, err
	}
	if user.ID != 0 {
		return nil, nil
	}
	newUser := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err = db.Create(&newUser).Error
	return &newUser, err
}

func UpdateUser(db *gorm.DB, name string, email string, password string) (*User, error) {
	var user User
	err := db.First(&user, "email = ?", email).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}

	user.Name = name
	user.Password = password

	err = db.Save(&user).Error
	return &user, err
}
