package keeper

import "github.com/carp-cobain/gin-todos/domain"

// TaskKeeper reads, writes, and deletes tasks
type TaskKeeper interface {
	TaskReader
	TaskWriter
	TaskDeleter
}

// TaskReader reads tasks
type TaskReader interface {
	// GetTask reads a single task
	GetTask(id uint64) (domain.Task, error)
	// GetTasks reads a page of tasks for a story
	GetTasks(storyID, cursor uint64, limit int) (uint64, []domain.Task)
}

// TaskWriter writes tasks
type TaskWriter interface {
	// CreateTask creates a new task
	CreateTask(storyID uint64, title string) (domain.Task, error)
	// UpdateTask updates an existing task
	UpdateTask(id uint64, title, status string) (domain.Task, error)
}

// TaskDeleter deletes tasks
type TaskDeleter interface {
	// DeleteTask deletes an existing task
	DeleteTask(id uint64) error
}
