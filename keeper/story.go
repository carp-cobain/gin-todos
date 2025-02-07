package keeper

import "github.com/carp-cobain/gin-todos/domain"

// StoryKeeper reads, writes, and deletes stories
type StoryKeeper interface {
	StoryReader
	StoryWriter
	StoryDeleter
}

// StoryReader reads stories
type StoryReader interface {
	// GetStory reads a single story
	GetStory(id uint64) (domain.Story, error)
	// GetStories reads a page of stories
	GetStories(cursor uint64, limit int) (uint64, []domain.Story)
}

// StoryWriter writes stories
type StoryWriter interface {
	// CreateStory creates a new story
	CreateStory(title string) (domain.Story, error)
	// UpdateStory updates an existing story
	UpdateStory(id uint64, title string) (domain.Story, error)
}

// StoryDeleter deletes stories
type StoryDeleter interface {
	// DeleteStory deletes an existing story
	DeleteStory(id uint64) error
}
