package models

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	BookName *string `json:"Book Name" gorm:"not_null"`
	Course   *string `json:"Course" gorm:"not_null"`
	Author   *string `json:"Author"`
	Sno      int64   `json:"Serial Number" gorm:"not_null;unique"`
}
