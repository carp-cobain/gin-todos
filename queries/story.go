package queries

import (
	"time"

	"github.com/carp-cobain/gin-todos/models"
)

// SelectStory selects a single story
func SelectStory(id uint64) (story models.Story, err error) {
	err = models.DB.Where("id = ?", id).First(&story).Error
	return
}

// SelectStories selects a page of stories
func SelectStories(limit, offset int) (stories []models.Story) {
	models.DB.Order("created_at").Limit(limit).Offset(offset).Find(&stories)
	return
}

// InsertStory inserts a new story
func InsertStory(title string) (story models.Story, err error) {
	story = models.Story{Title: title}
	err = models.DB.Create(&story).Error
	return
}

// UpdateStory updates a story
func UpdateStory(id uint64, title string) (story models.Story, err error) {
	if story, err = SelectStory(id); err == nil {
		err = models.DB.
			Model(&story).
			Updates(map[string]any{"title": title, "updated_at": time.Now()}).
			Error
	}
	return
}

// DeleteStory deletes a story
func DeleteStory(id uint64) (rows int64, err error) {
	var story models.Story
	if story, err = SelectStory(id); err == nil {
		result := models.DB.Delete(&story)
		rows, err = result.RowsAffected, result.Error
	}
	return
}
