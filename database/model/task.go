package model

import (
	"time"

	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

type Task struct {
	ID        uint64 `gorm:"primarykey"`
	StoryID   uint64 `gorm:"index"`
	Story     Story  `gorm:"foreignKey:StoryID"`
	Title     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (self Task) ToDomain() domain.Task {
	return domain.Task{
		ID:        self.ID,
		StoryID:   self.StoryID,
		Title:     self.Title,
		Status:    self.Status,
		CreatedAt: self.CreatedAt,
		UpdatedAt: self.UpdatedAt,
	}
}
