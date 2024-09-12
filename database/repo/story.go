package repo

import (
	"github.com/carp-cobain/gin-todos/database/model"
	"github.com/carp-cobain/gin-todos/database/query"
	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

// StoryRepo is a database keeper (reader and writer) of stories.
type StoryRepo struct {
	db *gorm.DB
}

// NewStoryRepo creates a new story repo
func NewStoryRepo(db *gorm.DB) StoryRepo {
	return StoryRepo{db}
}

// GetStory reads a single story from a database.
func (self StoryRepo) GetStory(id uint) (story domain.Story, err error) {
	var model model.Story
	if model, err = query.SelectStory(self.db, id); err == nil {
		story = model.ToDomain()
	}
	return
}

// GetStories reads a page of stories from a database.
func (self StoryRepo) GetStories(limit, offset int) []domain.Story {
	model := query.SelectStories(self.db, limit, offset)
	stories := make([]domain.Story, len(model))
	for i := 0; i < len(model); i++ {
		stories[i] = model[i].ToDomain()
	}
	return stories
}

// CreateStory inserts a new story into a database.
func (self StoryRepo) CreateStory(title string) (story domain.Story, err error) {
	var model model.Story
	if model, err = query.InsertStory(self.db, title); err == nil {
		story = model.ToDomain()
	}
	return
}

// UpdateStory updates a story in a database.
func (self StoryRepo) UpdateStory(id uint, title string) (story domain.Story, err error) {
	var model model.Story
	if model, err = query.UpdateStory(self.db, id, title); err == nil {
		story = model.ToDomain()
	}
	return
}

// RemoveStory deletes a story from a database.
func (self StoryRepo) DeleteStory(id uint) (int64, error) {
	return query.DeleteStory(self.db, id)
}
