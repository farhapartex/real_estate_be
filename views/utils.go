package views

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	return page, pageSize
}
