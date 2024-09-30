package repo

import (
	"github.com/carp-cobain/gin-todos/database/model"
	"github.com/carp-cobain/gin-todos/database/query"
	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

// StoryRepo is a database keeper of stories.
type StoryRepo struct {
	reader *gorm.DB
	writer *gorm.DB
}

// NewStoryRepo creates a new story repo
func NewStoryRepo(reader, writer *gorm.DB) StoryRepo {
	return StoryRepo{reader, writer}
}

// GetStory reads a single story from a database.
func (self StoryRepo) GetStory(id uint64) (story domain.Story, err error) {
	var model model.Story
	if model, err = query.SelectStory(self.reader, id); err == nil {
		story = model.ToDomain()
	}
	return
}

// GetStories reads a page of stories from a database.
func (self StoryRepo) GetStories(cursor uint64, limit int) (uint64, []domain.Story) {
	models := query.SelectStories(self.reader, cursor, limit)
	stories := make([]domain.Story, len(models))
	var nextCursor uint64
	for i, model := range models {
		stories[i] = model.ToDomain()
		if model.ID > nextCursor {
			nextCursor = model.ID
		}
	}
	return nextCursor, stories
}

// CreateStory inserts a story in a database.
func (self StoryRepo) CreateStory(title string) (story domain.Story, err error) {
	var model model.Story
	if model, err = query.InsertStory(self.writer, title); err == nil {
		story = model.ToDomain()
	}
	return
}

// UpdateStory updates a story in a database.
func (self StoryRepo) UpdateStory(id uint64, title string) (story domain.Story, err error) {
	var model model.Story
	if model, err = query.UpdateStory(self.writer, id, title); err == nil {
		story = model.ToDomain()
	}
	return
}

// RemoveStory deletes a story from a database.
func (self StoryRepo) DeleteStory(id uint64) error {
	return query.DeleteStory(self.writer, id)
}
