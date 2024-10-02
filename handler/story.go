package handler

import (
	"github.com/carp-cobain/gin-todos/keeper"
	"github.com/gin-gonic/gin"
)

// StoryHandler is the http/json api for managing stories
type StoryHandler struct {
	keeper keeper.StoryKeeper
}

// NewStoryHandler creates a new story handler
func NewStoryHandler(keeper keeper.StoryKeeper) StoryHandler {
	return StoryHandler{keeper}
}

// GET /stories
// Get a page of stories
func (self StoryHandler) GetStories(c *gin.Context) {
	cursor, limit := getPageParams(c)
	next, stories := self.keeper.GetStories(cursor, limit)
	okJson(c, gin.H{"cursor": next, "stories": stories})
}

// GET /stories/:id
// Get a story
func (self StoryHandler) GetStory(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, err)
		return
	}
	story, err := self.keeper.GetStory(id)
	if err != nil {
		notFound(c, err)
		return
	}
	okJson(c, gin.H{"story": story})
}

// POST /stories
// Create a new story
func (self StoryHandler) CreateStory(c *gin.Context) {
	var request StoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}
	title, err := request.Validate()
	if err != nil {
		badRequest(c, err)
		return
	}
	story, err := self.keeper.CreateStory(title)
	if err != nil {
		internalError(c, err)
		return
	}
	createdJson(c, gin.H{"story": story})
}

// PATCH /stories/:id
// Update a story
func (self StoryHandler) UpdateStory(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, err)
		return
	}
	var request StoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}
	title, err := request.Validate()
	if err != nil {
		badRequest(c, err)
		return
	}
	story, err := self.keeper.UpdateStory(id, title)
	if err != nil {
		notFound(c, err)
		return
	}
	okJson(c, gin.H{"story": story})
}

// DELETE /stories/:id
// Delete a story
func (self StoryHandler) DeleteStory(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, nil)
		return
	}
	if err := self.keeper.DeleteStory(id); err != nil {
		notFound(c, nil)
		return
	}
	noContent(c)
}
