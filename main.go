package main

import (
	"gin-webcore/routers"
)

func main() {

	router := routers.InitRouter()
	router.Run(":1002")
}
