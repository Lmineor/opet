package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	AllGood       = http.StatusOK
	NotFound      = http.StatusNotFound
	InternalError = http.StatusInternalServerError
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, Response{
		Data: data,
		Msg:  msg,
	})
}

func Success(data interface{}, message string, c *gin.Context) {
	Result(AllGood, data, message, c)
}

func FailedResp(c *gin.Context) {
	Result(NotFound, nil, "not found", c)
}

func ErrorResp(c *gin.Context) {
	Result(InternalError, nil, "internal error", c)
}