package controller

import (
	"DGoDoBackend/src/global"
	"DGoDoBackend/src/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/relvacode/iso8601"
)

func GetAllDDL(context *gin.Context) {
	name, _ := context.Cookie("user_name")

	ddls, err := model.GetAllDDLByName(name)

	if err != nil {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusInternalServerError, Msg: "服务端错误，查询失败"})
		return
	}

	context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusOK, Msg: "查询成功", Data: ddls})
}

func UpdateDDLStatus(context *gin.Context) {
	name, _ := context.Cookie("user_name")
	id, _ := strconv.Atoi(context.Query("id"))
	status, _ := strconv.Atoi(context.Query("status"))
	ddl, err := model.GetDDLByID(id)

	fmt.Println(status)
	fmt.Println(context.Query("status"))

	if err != nil {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusNotFound, Msg: "未找到相应DDL"})
		return
	}

	if ddl.Name != name {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusForbidden, Msg: "没有访问权限！"})
		return
	}

	ddl, err = ddl.UpdateDDLStatus(status)

	if err != nil {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusInternalServerError, Msg: "服务端错误，修改失败"})
		return
	}

	context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusOK, Msg: "修改成功", Data: ddl})
}

func CreateDDL(context *gin.Context) {
	name, _ := context.Cookie("user_name")
	content := context.Query("content")

	expect_time, err := iso8601.ParseString(context.Query("expect_time"))

	if err != nil {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusBadRequest, Msg: "参数错误"})
	}

	ddl := model.DDL{
		Name:        name,
		Content:     content,
		Create_time: time.Now(),
		Expect_time: expect_time,
		Status:      0,
	}

	ddl, err = ddl.CreateDDL()

	if err != nil {
		context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusInternalServerError, Msg: "服务端错误，新建失败"})
		return
	}

	context.JSON(http.StatusOK, global.ReturnType{Status: http.StatusOK, Msg: "创建成功", Data: ddl})

}
