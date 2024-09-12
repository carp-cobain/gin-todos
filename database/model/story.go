package model

import (
	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	Title string
}

func (self Story) ToDomain() domain.Story {
	return domain.Story{
		ID:        self.ID,
		Title:     self.Title,
		CreatedAt: self.CreatedAt,
		UpdatedAt: self.UpdatedAt,
	}
}
