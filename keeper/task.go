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
	GetTask(id uint64) (domain.Task, error)
	// GetTasks reads a page of tasks for a story
	GetTasks(storyID uint64, limit, offset int) []domain.Task
}

// TaskWriter writes tasks
type TaskWriter interface {
	// CreateTask creates a new story
	CreateTask(storyID uint64, name string) (domain.Task, error)
	// UpdateTask updates an existing story
	UpdateTask(id uint64, name, status string) (domain.Task, error)
	// DeleteTask deletes an existing story
	DeleteTask(id uint64) error
}
