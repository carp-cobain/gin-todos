package keeper

import "github.com/carp-cobain/gin-todos/domain"

// TaskKeeper reads and writes tasks
type TaskKeeper interface {
	TaskReader
	TaskWriter
}

// TaskReader reads tasks
type TaskReader interface {
	// GetTask reads a single task
	GetTask(id uint) (domain.Task, error)
	// GetTasks reads a page of tasks for a story
	GetTasks(storyID uint, limit, offset int) []domain.Task
}

// TaskWriter writes tasks
type TaskWriter interface {
	// CreateTask creates a new story
	CreateTask(storyID uint, name string) (domain.Task, error)
	// UpdateTask updates an existing story
	UpdateTask(id uint, name, status string) (domain.Task, error)
	// DeleteTask deletes an existing story
	DeleteTask(id uint) (int64, error)
}
