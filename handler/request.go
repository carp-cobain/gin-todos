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

// CreateTaskRequest is the request type for creating tasks
type CreateTaskRequest struct {
	StoryID uint64 `json:"storyId" binding:"required"`
	Name    string `json:"name" binding:"required,max=100"`
}

// UpdateTaskRequest is the request type for updating tasks
type UpdateTaskRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Check status
func (self UpdateTaskRequest) Validate() (string, string, error) {
	name := strings.TrimSpace(self.Name)
	status := strings.ToLower(strings.TrimSpace(self.Status))
	if name == "" && status == "" {
		return "", "", fmt.Errorf("no task update data provided")
	}
	if status != "" && status != "complete" && status != "incomplete" {
		return "", "", fmt.Errorf("status: invalid variant: %s", self.Status)
	}
	return name, status, nil
}

// Get and return bounded query parameters for paging. If no query params are found, default values
// are returned.
func getPageParams(c *gin.Context) (limit int, offset int) {
	limit, offset = 100, 0
	if limitQuery, ok := c.GetQuery("limit"); ok {
		limit, _ = strconv.Atoi(limitQuery)
	}
	if offsetQuery, ok := c.GetQuery("offset"); ok {
		offset, _ = strconv.Atoi(offsetQuery)
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	if offset < 0 {
		offset = 0
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
