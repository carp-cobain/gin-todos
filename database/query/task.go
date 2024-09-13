package query

import (
	"time"

	"github.com/carp-cobain/gin-todos/database/model"
	"gorm.io/gorm"
)

// SelectTask selects a single task from a database
func SelectTask(db *gorm.DB, id uint64) (task model.Task, err error) {
	err = db.Where("id = ?", id).First(&task).Error
	return
}

// SelectTasks selects a page of tasks for a task from a database
func SelectTasks(db *gorm.DB, storyID uint64, limit, offset int) (tasks []model.Task) {
	db.Where("story_id = ?", storyID).Order("created_at").Limit(limit).Offset(offset).
		Find(&tasks)
	return
}

// InsertTask inserts a new task into a database
func InsertTask(db *gorm.DB, storyID uint64, name string) (task model.Task, err error) {
	task = model.Task{StoryID: storyID, Name: name, Status: "incomplete"}
	err = db.Create(&task).Error
	return
}

// UpdateTask updates a task in a database
func UpdateTask(db *gorm.DB, id uint64, name, status string) (task model.Task, err error) {
	task, err = SelectTask(db, id)
	if err != nil {
		return
	}
	if name == "" {
		name = task.Name
	}
	if status == "" {
		status = task.Status
	}
	result := db.Model(&task).Updates(
		map[string]any{
			"name":       name,
			"status":     status,
			"updated_at": time.Now(),
		},
	)
	err = result.Error
	return
}

// DeleteTask deletes a task from a database
func DeleteTask(db *gorm.DB, id uint64) (rows int64, err error) {
	var task model.Task
	if task, err = SelectTask(db, id); err == nil {
		result := db.Delete(&task)
		rows, err = result.RowsAffected, result.Error
	}
	return
}
