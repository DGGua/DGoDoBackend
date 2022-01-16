package main

import (
	"DGoDoBackend/src/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

var httpServer *gin.Engine

func main() {
	// 开启mysql连接
	fmt.Println("hello world")
	// db, err := db.MySqlDb.DB()
	// if err != nil {
	// 	log.Print(err)
	// }
	// defer func() {
	// 	db.Close()
	// }()

	httpServer = gin.Default()

	server.Run(httpServer)
}
