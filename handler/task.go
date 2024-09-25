package handler

import (
	"github.com/carp-cobain/gin-todos/keeper"
	"github.com/gin-gonic/gin"
)

// TaskHandler is the http/json api for managing tasks
type TaskHandler struct {
	keeper keeper.TaskKeeper
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(keeper keeper.TaskKeeper) TaskHandler {
	return TaskHandler{keeper}
}

// GET /stories/:id/tasks
// Get a page of tasks for a story
func (self TaskHandler) GetTasks(c *gin.Context) {
	storyID, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, err)
		return
	}
	cursor, limit := getPageParams(c)
	next, tasks := self.keeper.GetTasks(storyID, cursor, limit)
	okJson(c, gin.H{"tasks": tasks, "cursor": next})
}

// GET /tasks/:id
// Get a task
func (self TaskHandler) GetTask(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, err)
		return
	}
	task, err := self.keeper.GetTask(id)
	if err != nil {
		notFound(c, err)
		return
	}
	okJson(c, gin.H{"task": task})

}

// POST /tasks
// Create a new task
func (self TaskHandler) CreateTask(c *gin.Context) {
	var request CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}
	storyID, title, err := request.Validate()
	if err != nil {
		badRequest(c, err)
		return
	}
	task, err := self.keeper.CreateTask(storyID, title)
	if err != nil {
		internalError(c, err)
		return
	}
	createdJson(c, gin.H{"task": task})
}

// PATCH /tasks/:id
// Update a task
func (self TaskHandler) UpdateTask(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, err)
		return
	}
	var request UpdateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}
	title, status, err := request.Validate()
	if err != nil {
		badRequest(c, err)
		return
	}
	task, err := self.keeper.UpdateTask(id, title, status)
	if err != nil {
		notFound(c, err)
		return
	}
	okJson(c, gin.H{"task": task})
}

// DELETE /tasks/:id
// Delete a task
func (self TaskHandler) DeleteTask(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, nil)
		return
	}
	if err := self.keeper.DeleteTask(id); err != nil {
		notFound(c, nil)
		return
	}
	noContent(c)
}
