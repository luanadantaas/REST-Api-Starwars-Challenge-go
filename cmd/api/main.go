package main

import (
	"starwars/src/database"
	"starwars/src/logger"
	"starwars/src/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo, err := database.New()
	if err != nil {
		logger.Log().Panic("couldn't connect to database: %v", err)
	}

	r := gin.Default()
	serve := service.New(repo, r)

	serve.Run()
}
