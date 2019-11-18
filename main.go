package main

import (
	"fmt"
	"gin-webcore/database"
	"gin-webcore/routers"
)

func main() {
	fmt.Println("main.go")
	defer database.DB.Close()

	router := routers.InitRouter()
	router.Run(":1002")
}
