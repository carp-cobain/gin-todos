package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func uintParam(c *gin.Context, key string) (i uint64, err error) {
	value := c.Param(key)
	if i, err = strconv.ParseUint(value, 10, 64); err != nil {
		err = fmt.Errorf("%s: expected uint64, got: %s", key, value)
	}
	return
}
