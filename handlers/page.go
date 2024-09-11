package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Max page size
const maxLimit int = 1000

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
