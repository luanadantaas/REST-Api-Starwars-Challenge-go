package service

import (
	"net/http"
	"starwars/src/entity"
	"starwars/src/logger"
	"starwars/src/repository"
	"starwars/src/swapi"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service struct {
	router     *gin.Engine
	repository repository.Repository
}

func New(repo repository.Repository, router *gin.Engine) *Service {
	return &Service{
		router:     router,
		repository: repo,
	}
}

func (s *Service) Run() {
	//list all planets
	s.router.GET("/planets", s.AllPlanets)
	//get planet by name
	s.router.GET("/planets/:name", s.FindByName)
	//get planet by id
	s.router.GET("/planets/:id", s.FindById)
	//add planet
	s.router.POST("/planets/", s.Add)
	//delete planet
	s.router.DELETE("/planets/:id", s.Delete)

	s.router.Run() // listen and serve on 0.0.0.0:8080
}

func (s *Service) AllPlanets(c *gin.Context) {
	list, err := s.repository.ListPlanet()
	if err != nil {
		logger.Log().Warn("%v", err)
	}

	c.JSON(http.StatusOK, list)
}

func (s *Service) FindByName(c *gin.Context) {
	name := c.Param("name")
	findingName, err := s.repository.FindByName(name)
	if err != nil {
		logger.Log().Warn("%v", err)
	}

	c.JSON(http.StatusOK, findingName)
}

func (s *Service) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log().Warn("%v", err)
	}

	findingId, err := s.repository.FindById(id)
	if err != nil {
		logger.Log().Warn("%v", err)
	}

	c.JSON(http.StatusOK, findingId)
}

//função que vai interagir com a outra api
func (s *Service) Add(c *gin.Context) {
	var planet entity.Planet
	c.Bind(&planet)

	planet.Appearances = swapi.GetAppearances(planet.Name)
	err := s.repository.AddPlanet(planet)
	if err != nil {
		logger.Log().Warn("%v", err)
	}
	c.JSON(http.StatusOK, err)
}

func (s *Service) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log().Warn("%v", err)
	}

	err = s.repository.RemovePlanet(id)
	if err != nil {
		logger.Log().Warn("%v", err)
	}
	c.JSON(http.StatusOK, err)
}
