package handlers

import (
	"strings"

	"github.com/carp-cobain/gin-todos/dtos"
	"github.com/carp-cobain/gin-todos/queries"

	"github.com/gin-gonic/gin"
)

// GET /stories
// List a page of stories
func ListStories(c *gin.Context) {
	limit, offset := getPageParams(c)
	stories := queries.SelectStories(limit, offset)
	okJson(c, gin.H{
		"limit":   limit,
		"offset":  offset,
		"stories": dtos.NewStoriesResponse(stories),
	})
}

// GET /stories/:id
// Get a story
func GetStory(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, err)
		return
	}
	story, err := queries.SelectStory(id)
	if err != nil {
		notFound(c, err)
		return
	}
	okJson(c, gin.H{
		"story": dtos.NewStoryResponse(story),
	})
}

// POST /stories
// Create a new story
func CreateStory(c *gin.Context) {
	var request dtos.StoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}
	story, err := queries.InsertStory(strings.TrimSpace(request.Title))
	if err != nil {
		internalError(c, err)
		return
	}
	okJson(c, gin.H{
		"story": dtos.NewStoryResponse(story),
	})
}

// PATCH /stories/:id
// Update a story
func UpdateStory(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, err)
		return
	}
	var request dtos.StoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}
	story, err := queries.UpdateStory(id, strings.TrimSpace(request.Title))
	if err != nil {
		notFound(c, err)
		return
	}
	okJson(c, gin.H{
		"story": dtos.NewStoryResponse(story),
	})
}

// DELETE /stories/:id
// Delete a story
func DeleteStory(c *gin.Context) {
	id, err := uintParam(c, "id")
	if err != nil {
		badRequest(c, nil)
		return
	}
	if _, err := queries.DeleteStory(id); err != nil {
		notFound(c, nil)
		return
	}
	noContent(c)
}
