package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"infinite-window-micro/pkg/setting"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}

// GetOffset
func GetOffset(page, pageSize int) int {
	offset := 0

	if pageSize <= 0 {
		pageSize = setting.AppSetting.PageSize
	}
	if page > 0 {
		offset = (page - 1) * pageSize
	}

	return offset
}
