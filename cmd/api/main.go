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

	serve := service.New(repo)
	r := gin.Default()
	//list all planets
	r.GET("/planets", serve.AllPlanets)
	//get planet by name
	r.GET("/planets/:name", serve.FindByName)
	//get planet by id
	r.GET("/planets/:id", serve.FindById)
	//add planet
	r.POST("/planets/", serve.Add)
	//delete planet
	r.DELETE("/planets/:id", serve.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
