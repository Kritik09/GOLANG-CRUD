package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      *string `json:"name" gorm:"not null"`
	Email     *string `json:"email" gorm:"unique;not null"`
	Course    *string `json:"course" gorm:"not null"`
	MobNumber int64   `json:"Phone Number" gorm:"not null;unique"`
}
