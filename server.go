package main

import (
	"awesomeProject/config"
	"awesomeProject/db"
	"awesomeProject/router"
	"flag"
	"fmt"
)

func main() {
	setConfig()

	db.Init()

	c := config.Config.Database
	fmt.Printf("DBユーザー::%s", c.User)

	e := router.Init()
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