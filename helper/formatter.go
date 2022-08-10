package helper

import (
	"github.com/gin-gonic/gin"
)

type ReturnType struct {
	Status int
	Msg    string
	Data   interface{}
}

func ApiReturn(status int, msg string, data interface{}) gin.H {
	return gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	}
}
