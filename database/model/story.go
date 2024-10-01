package model

import (
	"github.com/carp-cobain/gin-todos/domain"
)

type Story struct {
	ID        uint64 `gorm:"primarykey"`
	Title     string
	CreatedAt Time
	UpdatedAt Time
}

func (self Story) ToDomain() domain.Story {
	return domain.Story{
		ID:        self.ID,
		Title:     self.Title,
		CreatedAt: self.CreatedAt.FromUnix(),
		UpdatedAt: self.UpdatedAt.FromUnix(),
	}
}
