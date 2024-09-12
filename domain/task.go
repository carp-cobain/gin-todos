package domain

import "time"

// Task is a single unit of work under a story.
type Task struct {
	ID        uint      `json:"id"`
	StoryID   uint      `json:"storyId"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
