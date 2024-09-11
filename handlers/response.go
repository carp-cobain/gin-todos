package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func okJson(c *gin.Context, body gin.H) {
	c.JSON(http.StatusOK, body)
}

func errorResponse(c *gin.Context, status int, err error) {
	if err == nil {
		c.Status(status)
		return
	}
	c.JSON(status, gin.H{"error": err.Error()})
}

func badRequest(c *gin.Context, err error) {
	errorResponse(c, http.StatusBadRequest, err)
}

func notFound(c *gin.Context, err error) {
	errorResponse(c, http.StatusNotFound, err)
}

func noContent(c *gin.Context) {
	errorResponse(c, http.StatusNoContent, nil)
}

func internalError(c *gin.Context, err error) {
	log.Printf("INTERNAL ERROR: %+v", err)
	errorResponse(c, http.StatusInternalServerError, err)
}
