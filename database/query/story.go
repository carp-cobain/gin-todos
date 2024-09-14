package query

import (
	"time"

	"github.com/carp-cobain/gin-todos/database/model"
	"gorm.io/gorm"
)

// SelectStory selects a single story from a database
func SelectStory(db *gorm.DB, id uint64) (story model.Story, err error) {
	err = db.Where("id = ?", id).First(&story).Error
	return
}

// SelectStories selects a page of stories from a database
func SelectStories(db *gorm.DB, limit, offset int) (stories []model.Story) {
	db.Order("created_at").Limit(limit).Offset(offset).Find(&stories)
	return
}

// InsertStory inserts a new story into a database
func InsertStory(db *gorm.DB, title string) (story model.Story, err error) {
	story = model.Story{Title: title}
	err = db.Create(&story).Error
	return
}

// UpdateStory updates a story in a database
func UpdateStory(db *gorm.DB, id uint64, title string) (story model.Story, err error) {
	if story, err = SelectStory(db, id); err == nil {
		err = db.
			Model(&story).
			Updates(updates{"title": title, "updated_at": time.Now()}).
			Error
	}
	return
}

// DeleteStory deletes a story from a database
func DeleteStory(db *gorm.DB, id uint64) (rows int64, err error) {
	var story model.Story
	if story, err = SelectStory(db, id); err == nil {
		result := db.Delete(&story)
		rows, err = result.RowsAffected, result.Error
	}
	return
}
