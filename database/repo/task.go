package repo

import (
	"github.com/carp-cobain/gin-todos/database/model"
	"github.com/carp-cobain/gin-todos/database/query"
	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

// TaskRepo is a database keeper (reader and writer) of tasks.
type TaskRepo struct {
	db *gorm.DB
}

// NewTaskRepo creates a new task repo
func NewTaskRepo(db *gorm.DB) TaskRepo {
	return TaskRepo{db}
}

// GetTask reads a single story from a database.
func (self TaskRepo) GetTask(id uint64) (task domain.Task, err error) {
	var model model.Task
	if model, err = query.SelectTask(self.db, id); err == nil {
		task = model.ToDomain()
	}
	return
}

// GetTasks reads a page of tasks from a database.
func (self TaskRepo) GetTasks(storyID uint64, limit, offset int) []domain.Task {
	model := query.SelectTasks(self.db, uint64(storyID), limit, offset)
	tasks := make([]domain.Task, len(model))
	for i := 0; i < len(model); i++ {
		tasks[i] = model[i].ToDomain()
	}
	return tasks
}

// CreateTask inserts a new story into a database.
func (self TaskRepo) CreateTask(storyID uint64, name string) (story domain.Task, err error) {
	var model model.Task
	if model, err = query.InsertTask(self.db, uint64(storyID), name); err == nil {
		story = model.ToDomain()
	}
	return
}

// UpdateTask updates a story in a database.
func (self TaskRepo) UpdateTask(id uint64, name, status string) (story domain.Task, err error) {
	var model model.Task
	if model, err = query.UpdateTask(self.db, id, name, status); err == nil {
		story = model.ToDomain()
	}
	return
}

// DeleteTask deletes a story from a database.
func (self TaskRepo) DeleteTask(id uint64) (int64, error) {
	return query.DeleteTask(self.db, id)
}
