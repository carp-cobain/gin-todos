package dtos

import (
	"time"

	"github.com/carp-cobain/gin-todos/models"
)

// StoryRequest is the request type for creating and updating stories.
type StoryRequest struct {
	Title string `json:"title" binding:"required,max=100"`
}

// StoryResponse is the data trasnfer object for http responses.
type StoryResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewStoriesResponse creates a slice of story response data transfer objects
func NewStoriesResponse(stories []models.Story) []StoryResponse {
	response := make([]StoryResponse, len(stories))
	for i := 0; i < len(stories); i++ {
		response[i] = NewStoryResponse(stories[i])
	}
	return response
}

// NewStoryResponse creates a new story data transfer object
func NewStoryResponse(story models.Story) StoryResponse {
	return StoryResponse{
		ID:        story.ID,
		Title:     story.Title,
		CreatedAt: story.CreatedAt,
		UpdatedAt: story.UpdatedAt,
	}
}
