package controller

import (
	"DGoDoBackend/src/global"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetHistoryFile(context *gin.Context) {
	const filename = "docs/history.md"
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusInternalServerError, Msg: "服务端错误，未找到文件"})
		return
	}
	context.Data(http.StatusOK, "text/html", fileContent)

}

func GetTodoFile(context *gin.Context) {
	const filename = "docs/todo.md"
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusInternalServerError, Msg: "服务端错误，未找到文件"})
		return
	}
	context.Data(http.StatusOK, "text/html", fileContent)

}
