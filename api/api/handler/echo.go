package handler

import "github.com/gin-gonic/gin"

func Echo(c *gin.Context) {
	c.Writer.Write([]byte("api echo."))
}
