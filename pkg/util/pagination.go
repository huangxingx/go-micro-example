package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"go-micro-example/pkg/setting"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int32 {
	var result int32
	page := int32(com.StrTo(c.Query("page")).MustInt())
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}

// GetOffset
func GetOffset(page, pageSize int32) int32 {
	var offset int32

	if pageSize <= 0 {
		pageSize = setting.AppSetting.PageSize
	}
	if page > 0 {
		offset = (page - 1) * pageSize
	}

	return offset
}
