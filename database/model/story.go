package model

import (
	"time"

	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

type Story struct {
	ID        uint64 `gorm:"primarykey"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (self Story) ToDomain() domain.Story {
	return domain.Story{
		ID:        self.ID,
		Title:     self.Title,
		CreatedAt: self.CreatedAt,
		UpdatedAt: self.UpdatedAt,
	}
}
