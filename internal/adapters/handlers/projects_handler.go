package handlers

import (
	"my-stats/internal/domain"
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

type SendProjects struct {
	Result []domain.Project `json:"result"`
}

func (h *ProjectsHandler) GetProjects(c *gin.Context) {
	projects := SendProjects{
		Result: h.service.Get(),
	}

	c.JSON(http.StatusOK, projects)
}
