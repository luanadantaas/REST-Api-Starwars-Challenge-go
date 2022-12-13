package service

import (
	"net/http"
	"starwars/src/entity"
	"starwars/src/logger"
	"starwars/src/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service struct {
	router     gin.IRouter
	repository repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{
		router:     &gin.RouterGroup{},
		repository: repo,
	}
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
