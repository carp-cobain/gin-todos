package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Max page size
const maxLimit int = 1000

// StoryRequest is the request type for creating and updating stories.
type StoryRequest struct {
	Title string `json:"title" binding:"required,max=100"`
}

// Validate request params
func (self StoryRequest) Validate() (string, error) {
	title := strings.TrimSpace(self.Title)
	if title == "" {
		return "", fmt.Errorf("story title cannot be blank")
	}
	return title, nil
}

// CreateTaskRequest is the request type for creating tasks
type CreateTaskRequest struct {
	StoryID uint64 `json:"storyId" binding:"required,min=1"`
	Title   string `json:"title" binding:"required,max=100"`
}

// Validate request params
func (self CreateTaskRequest) Validate() (uint64, string, error) {
	title := strings.TrimSpace(self.Title)
	if title == "" {
		return 0, "", fmt.Errorf("task title cannot be blank")
	}
	return self.StoryID, title, nil
}

// UpdateTaskRequest is the request type for updating tasks
type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

// Check status
func (self UpdateTaskRequest) Validate() (string, string, error) {
	title := strings.TrimSpace(self.Title)
	status := strings.ToLower(strings.TrimSpace(self.Status))
	if title == "" && status == "" {
		return "", "", fmt.Errorf("no task update provided")
	}
	if status != "" && status != "complete" && status != "incomplete" {
		return "", "", fmt.Errorf("status: invalid variant: %s", self.Status)
	}
	return title, status, nil
}

// Get and return bounded query parameters for paging. If no query params are found, default values
// are returned.
func getPageParams(c *gin.Context) (cursor uint64, limit int) {
	cursor, limit = 0, 100
	if limitQuery, ok := c.GetQuery("limit"); ok {
		limit, _ = strconv.Atoi(limitQuery)
	}
	if cursorQuery, ok := c.GetQuery("cursor"); ok {
		cursor, _ = strconv.ParseUint(cursorQuery, 10, 64)
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	if cursor < 0 {
		cursor = 0
	}
	return
}

// Read an unsigned integer parameter with the given key
func uintParam(c *gin.Context, key string) (uint64, error) {
	value := c.Param(key)
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: expected uint64, got: %s", key, value)
	}
	return i, nil
}
