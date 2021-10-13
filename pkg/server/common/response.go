package common

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int,data interface{}, msg string, c *gin.Context){
	c.JSON(code, Response{
		Data: data,
		Msg: msg,
	})
}

func Success(data interface{}, message string, c *gin.Context){
	Result(AllGood, data, message, c)
}
