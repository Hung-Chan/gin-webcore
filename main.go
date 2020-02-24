package main

import (
	"fmt"
	"gin-webcore/database"
	_ "gin-webcore/docs"
	"gin-webcore/routers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	// autoload .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loading .env file")
	}

	// init Database
	drive := os.Getenv("DB_CONNECTION")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbDatabase)

	database.Conn(drive, conn)

	// init Redis
}

// @title Golang Gin-Webcore API
// @version 1.0
// @description This is a Gin-webcore
// @host localhost:1002
func main() {
	fmt.Println("main.go")
	defer database.DB.Close()

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
