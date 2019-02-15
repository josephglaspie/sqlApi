package main

import (
	"github.com/sqlApi/app"
	"github.com/sqlApi/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
