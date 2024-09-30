package repo

import (
	"github.com/carp-cobain/gin-todos/database/model"
	"github.com/carp-cobain/gin-todos/database/query"
	"github.com/carp-cobain/gin-todos/domain"
	"gorm.io/gorm"
)

// TaskRepo is a database keeper of tasks.
type TaskRepo struct {
	reader *gorm.DB
	writer *gorm.DB
}

// NewTaskRepo creates a new task repo
func NewTaskRepo(reader, writer *gorm.DB) TaskRepo {
	return TaskRepo{reader, writer}
}

// GetTask reads a single task from a database.
func (self TaskRepo) GetTask(id uint64) (task domain.Task, err error) {
	var model model.Task
	if model, err = query.SelectTask(self.reader, id); err == nil {
		task = model.ToDomain()
	}
	return
}

// GetTasks reads a page of tasks from a database.
func (self TaskRepo) GetTasks(storyID, cursor uint64, limit int) (uint64, []domain.Task) {
	models := query.SelectTasks(self.reader, storyID, cursor, limit)
	tasks := make([]domain.Task, len(models))
	var nextCursor uint64
	for i, model := range models {
		tasks[i] = model.ToDomain()
		if model.ID > nextCursor {
			nextCursor = model.ID
		}
	}
	return nextCursor, tasks
}

// CreateTask inserts a task in a database.
func (self TaskRepo) CreateTask(storyID uint64, title string) (task domain.Task, err error) {
	var model model.Task
	if model, err = query.InsertTask(self.writer, storyID, title); err == nil {
		task = model.ToDomain()
	}
	return
}

// UpdateTask updates a task in a database.
func (self TaskRepo) UpdateTask(id uint64, title, status string) (task domain.Task, err error) {
	var model model.Task
	if model, err = query.UpdateTask(self.writer, id, title, status); err == nil {
		task = model.ToDomain()
	}
	return
}

// DeleteTask deletes a task from a database.
func (self TaskRepo) DeleteTask(id uint64) error {
	return query.DeleteTask(self.writer, id)
}
