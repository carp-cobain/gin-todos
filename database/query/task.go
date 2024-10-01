package query

import (
	"github.com/carp-cobain/gin-todos/database/model"
	"gorm.io/gorm"
)

// SelectTask selects a single task from a database
func SelectTask(db *gorm.DB, id uint64) (task model.Task, err error) {
	err = db.Where("id = ?", id).First(&task).Error
	return
}

// SelectTasks selects a page of tasks for a task from a database
func SelectTasks(db *gorm.DB, storyID, cursor uint64, limit int) (tasks []model.Task) {
	db.Where("story_id = ?", storyID).Where("id > ?", cursor).Order("id").Limit(limit).Find(&tasks)
	return
}

// InsertTask inserts a new task into a database
func InsertTask(db *gorm.DB, storyID uint64, title string) (task model.Task, err error) {
	task = model.Task{StoryID: storyID, Title: title, Status: "incomplete"}
	err = db.Create(&task).Error
	return
}

// UpdateTask updates a task in a database
func UpdateTask(db *gorm.DB, id uint64, title, status string) (task model.Task, err error) {
	task, err = SelectTask(db, id)
	if err != nil {
		return
	}
	if title == "" {
		title = task.Title
	}
	if status == "" {
		status = task.Status
	}
	result := db.Model(&task).Updates(updates{"title": title, "status": status})
	err = result.Error
	return
}

// DeleteTask deletes a task from a database
func DeleteTask(db *gorm.DB, id uint64) (err error) {
	var task model.Task
	if task, err = SelectTask(db, id); err == nil {
		err = db.Delete(&task).Error
	}
	return
}
