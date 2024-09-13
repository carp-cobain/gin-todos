package domain

import (
	"time"
)

// Story is something that needs to be done.
type Story struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
