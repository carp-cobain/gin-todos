package models

import "gorm.io/gorm"

type Story struct {
	gorm.Model
	Title string
}
