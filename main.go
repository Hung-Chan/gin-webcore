package main

import (
	"fmt"
	"gin-webcore/database"
	_ "gin-webcore/docs"
	"gin-webcore/routers"
)

// @title Golang Gin-Webcore API
// @version 1.0
// @description This is a Gin-webcore
// @host localhost:1002
func main() {
	fmt.Println("main.go")
	defer database.DB.Close()

	router := routers.InitRouter()
	router.Run(":1002")
}
