package app

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-micro-example/pkg/e"
	"go-micro-example/pkg/util"
)

type Gin struct {
	C *gin.Context
}

func NewAppGin(c *gin.Context) (appG Gin) {

	return Gin{C: c}
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponsePageData struct {
	DataList interface{} `json:"data_list"`
	Count    int         `json:"count"`
}

func (r ResponsePageData) GetMd5() (md5String string) {
	dataBytes, err := json.Marshal(r.DataList)
	if err != nil {
		return
	}
	md5String = util.EncodeMD5(string(dataBytes) + string(r.Count))

	return
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// ResponseWithMsg
func (g *Gin) ResponseWithMsg(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

// Response Success
func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	})
	return

}

// Response Error
func (g *Gin) Fail(errCode int, data interface{}) {

	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// FailWithErrorMsg Error msg
func (g *Gin) FailWithErrorMsg(errCode int, msg string, data interface{}) {
	httpCode := http.StatusOK

	if errCode == http.StatusBadRequest {
		httpCode = 400
	}

	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func (g *Gin) FailWithMsg(msg string, data interface{}) {

	g.C.JSON(http.StatusInternalServerError, Response{
		Code: e.ERROR,
		Msg:  msg,
		Data: data,
	})
	return
}
