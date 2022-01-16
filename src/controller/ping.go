package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "你好",
	})
}
