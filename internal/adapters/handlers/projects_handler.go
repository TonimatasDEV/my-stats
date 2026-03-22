package handlers

import (
	"my-stats/internal/ports/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectsHandler struct {
	service *services.ProjectsService
}

func NewProjectsHandler(service *services.ProjectsService) *ProjectsHandler {
	return &ProjectsHandler{service: service}
}

type SendStats struct {
	Names  []string `json:"names"`
	Values []int    `json:"values"`
}

func (h *ProjectsHandler) GetProjects(c *gin.Context) {
	var names []string
	var values []int

	for _, project := range h.service.Get() {
		names = append(names, project.Name)
		values = append(values, project.Downloads)
	}

	projects := SendStats{
		Names:  names,
		Values: values,
	}

	c.JSON(http.StatusOK, projects)
}
