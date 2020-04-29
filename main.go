package main

import (
	"awesomeProject/config"
	"awesomeProject/db"
	"awesomeProject/handler"
	"awesomeProject/log"
	"awesomeProject/router"
	"flag"
	"fmt"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1313
// @BasePath /api/v1
func main() {
	setConfig()

	db.Init()

	log.Init()

	c := config.Config.Database
	fmt.Printf("DBユーザー::%s", c.User)

	e := router.Init()
	e.HTTPErrorHandler = handler.JSONErrorHandler
	e.Logger.Fatal(e.Start(":8080"))
}

func setConfig() {
	env := "development"
	flag.Parse()
	if args := flag.Args(); 0 < len(args) && args[0] == "pro" {
		env = "production"
	}
	config.SetEnvironment(env)
}