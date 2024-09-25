package domain

import "time"

// Task is a single unit of work under a story.
type Task struct {
	ID        uint64    `json:"id"`
	StoryID   uint64    `json:"storyId"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
