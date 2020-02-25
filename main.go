package main

import (
	"fmt"
	"gin-webcore/database"
	_ "gin-webcore/docs"
	"gin-webcore/redis"
	"gin-webcore/routers"
	"net/http"
	"time"
)

// @title Golang Gin-Webcore API
// @version 1.0
// @description This is a Gin-webcore
// @host localhost:1002
func main() {
	fmt.Println("main.go")
	defer database.DB.Close()

	defer redis.RedisManage.Close()

	router := routers.InitRouter()

	server := &http.Server{
		Addr:           ":1002",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
