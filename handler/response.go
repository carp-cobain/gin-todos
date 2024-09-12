package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sends a 200 JSON response
func okJson(c *gin.Context, body gin.H) {
	c.JSON(http.StatusOK, body)
}

// Sends a 201 JSON response
func createdJson(c *gin.Context, body gin.H) {
	c.JSON(http.StatusCreated, body)
}

// Sends an error JSON response.
func errorResponse(c *gin.Context, status int, err error) {
	if err == nil {
		c.Status(status)
		return
	}
	c.JSON(status, gin.H{"error": err.Error()})
}

// Sends a 400 error JSON response.
func badRequest(c *gin.Context, err error) {
	errorResponse(c, http.StatusBadRequest, err)
}

// Sends a 404 error JSON response.
func notFound(c *gin.Context, err error) {
	errorResponse(c, http.StatusNotFound, err)
}

// Sends a 204 status.
func noContent(c *gin.Context) {
	errorResponse(c, http.StatusNoContent, nil)
}

// Sends a 500 error JSON response.
func internalError(c *gin.Context, err error) {
	log.Printf("INTERNAL ERROR: %+v", err)
	errorResponse(c, http.StatusInternalServerError, err)
}
