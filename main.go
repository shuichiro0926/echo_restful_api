package main

import (
	"awesomeProject/infrastructure"
	"awesomeProject/infrastructure/router"
)

func main() {
	infrastructure.DBInit()

	e := router.LoggerRouter()
	e.Logger.Fatal(e.Start(":8080"))
}