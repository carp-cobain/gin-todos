package model

import (
	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	StoryID uint
	Story   Story `gorm:"foreignKey:StoryID"`
	Name    string
	Status  string
}

func (self Task) ToDomain() domain.Task {
	return domain.Task{
		ID:        self.ID,
		StoryID:   self.StoryID,
		Name:      self.Name,
		Status:    self.Status,
		CreatedAt: self.CreatedAt,
		UpdatedAt: self.UpdatedAt,
	}
}
