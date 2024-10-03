package repo

import (
	"github.com/carp-cobain/gin-todos/database/model"
	"github.com/carp-cobain/gin-todos/database/query"
	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

// TaskRepo is a database keeper of tasks.
type TaskRepo struct {
	readDB  *gorm.DB
	writeDB *gorm.DB
}

// NewTaskRepo creates a new task repo
func NewTaskRepo(readDB, writeDB *gorm.DB) TaskRepo {
	return TaskRepo{readDB, writeDB}
}

// GetTask reads a single task from a database.
func (self TaskRepo) GetTask(id uint64) (task domain.Task, err error) {
	var model model.Task
	if model, err = query.SelectTask(self.readDB, id); err == nil {
		task = model.ToDomain()
	}
	return
}

// GetTasks reads a page of tasks from a database.
func (self TaskRepo) GetTasks(storyID, cursor uint64, limit int) (next uint64, tasks []domain.Task) {
	models := query.SelectTasks(self.readDB, storyID, cursor, limit)
	tasks = make([]domain.Task, len(models))
	for i, model := range models {
		tasks[i] = model.ToDomain()
		next = max(next, model.ID)
	}
	return
}

// CreateTask inserts a task in a database.
func (self TaskRepo) CreateTask(storyID uint64, title string) (task domain.Task, err error) {
	var model model.Task
	if model, err = query.InsertTask(self.writeDB, storyID, title); err == nil {
		task = model.ToDomain()
	}
	return
}

// UpdateTask updates a task in a database.
func (self TaskRepo) UpdateTask(id uint64, title, status string) (task domain.Task, err error) {
	var model model.Task
	if model, err = query.UpdateTask(self.writeDB, id, title, status); err == nil {
		task = model.ToDomain()
	}
	return
}

// DeleteTask deletes a task from a database.
func (self TaskRepo) DeleteTask(id uint64) error {
	return query.DeleteTask(self.writeDB, id)
}
