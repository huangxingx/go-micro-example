package app

import (
	"fmt"
	"go-micro-example/constant"
	"strings"

	"github.com/unknwon/com"
	"go-micro-example/pkg/e"
	"go-micro-example/pkg/setting"
)

type errorNoAuth struct {
	ErrorCode   int
	ErrorString string
}

func (e errorNoAuth) Error() string {
	return e.ErrorString
}

var ErrorNoAuth = errorNoAuth{e.NO_AUTH, e.GetMsg(e.NO_AUTH)}

func (g *Gin) GetAuthUserId() uint {
	userId, exists := g.C.Get("userId")
	if !exists {
		panic(ErrorNoAuth)
	}
	userIdInt := uint(userId.(int))
	return userIdInt
}

func (g *Gin) GetPage() int32 {
	page := com.StrTo(g.C.DefaultQuery("page", "0")).MustInt()
	if page <= 0 {
		page = 1
	}

	return int32(page)
}

func (g *Gin) GetPageSize() int32 {
	pageSize := int32(com.StrTo(g.C.DefaultQuery("page_size", "0")).MustInt())
	if pageSize <= 0 {
		pageSize = setting.AppSetting.PageSize
	}
	return pageSize
}

func (g *Gin) GetSearchKey() string {
	return g.C.DefaultQuery("search_key", "")

}

func (g *Gin) GetHostUrl() string {
	proto := "http"
	if strings.HasPrefix(g.C.Request.Proto, "HTTPS") {
		proto = "https"
	}

	return fmt.Sprintf("%s://%s", proto, g.C.Request.Host)

}

//获取app设备类型
func (g *Gin) GetAppPlatformType() uint8 {

	userAgent := strings.ToLower(g.C.Request.Header.Get("User-Agent"))
	if strings.Contains(userAgent, "android") {
		return constant.PLATFORM_ANDROID

	} else if strings.Contains(userAgent, "iphone") || strings.Contains(userAgent, "ipad") {
		return constant.PLATFORM_IOS
	}

	return constant.PLATFORM_OTHER
}
