package model

import (
	"github.com/carp-cobain/gin-todos/domain"
)

type Task struct {
	ID        uint64 `gorm:"primarykey"`
	StoryID   uint64 `gorm:"index"`
	Story     Story  `gorm:"foreignKey:StoryID"`
	Title     string
	Status    string
	CreatedAt Time
	UpdatedAt Time
}

func (self Task) ToDomain() domain.Task {
	return domain.Task{
		ID:        self.ID,
		StoryID:   self.StoryID,
		Title:     self.Title,
		Status:    self.Status,
		CreatedAt: self.CreatedAt.FromUnix(),
		UpdatedAt: self.UpdatedAt.FromUnix(),
	}
}
