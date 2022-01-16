package router

import (
	"DGoDoBackend/src/controller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/ping", controller.Ping)

	// DDLGroup := router.Group("/ddl")
	// DDLGroup.GET("/all", controller.GetAllDDL)
	// DDLGroup.POST("/updateStatus", controller.UpdateDDLStatus)
	// DDLGroup.POST("/create", controller.CreateDDL)

	DocGroup := router.Group("/docs")
	DocGroup.GET("/getHistory", controller.GetHistoryFile)
	DocGroup.GET("/getTodo", controller.GetTodoFile)
	// router.POST("/mail", Controller.RequestEmailValidation)

	// user := router.Group("/user")
	// {
	// 	user.POST("/register", Controller.Register)
	// 	user.POST("/login", Controller.DoLogin)
	// 	user.POST("/logout", Controller.DoLogout)
	// }

	// todo := router.Group("/todo")
	// todo.Use(middleware.Authority())
	// {
	// 	todo.POST("/add", Controller.AddTodo)
	// 	todo.GET("/list", Controller.GetUsersTodoList)
	// 	todo.PUT("/:id", Controller.UpdateTodoItem)
	// 	todo.DELETE("/:id", Controller.DeleteTodoItem)
	// }

	// TodoGroup := router.Group("/todo_group")
	// todo.Use(middleware.Authority())
	// {
	// 	TodoGroup.POST("/add", Controller.AddTodoGroup)
	// 	TodoGroup.GET("/list", Controller.GetAllTodoGroup)
	// 	TodoGroup.PUT("/:id", Controller.UpdateTodoGroup)
	// 	TodoGroup.DELETE("/:id", Controller.DeleteTodoGroup)
	// }

}
