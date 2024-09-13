package keeper

import "github.com/carp-cobain/gin-todos/domain"

// StoryKeeper reads and writes stories
type StoryKeeper interface {
	StoryReader
	StoryWriter
}

// StoryReader reads stories
type StoryReader interface {
	// GetStory reads a single story
	GetStory(id uint64) (domain.Story, error)
	// GetStories reads a page of stories
	GetStories(limit, offset int) []domain.Story
}

// StoryWriter writes stories
type StoryWriter interface {
	// CreateStory creates a new story
	CreateStory(title string) (domain.Story, error)
	// UpdateStory updates an existing story
	UpdateStory(id uint64, title string) (domain.Story, error)
	// DeleteStory deletes an existing story
	DeleteStory(id uint64) (int64, error)
}
