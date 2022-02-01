package main

import (
	"github.com/borischen0203/Go-todolist/config"
	"github.com/borischen0203/Go-todolist/logger"
	"github.com/borischen0203/Go-todolist/router"

	_ "github.com/joho/godotenv/autoload"
)

func Setup() {
	logger.Setup()
	config.Setup()
	router.Setup()
}

func main() {
	Setup()
	router.Router.Run()
}
