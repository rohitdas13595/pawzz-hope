package models

import (
	"gorm.io/gorm"
)

// gorm.Model definition
type Admin struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string
}

func (a *Admin) SignUp() {

}
