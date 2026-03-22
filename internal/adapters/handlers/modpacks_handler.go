package handlers

import (
	"my-stats/internal/ports/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModpacksHandler struct {
	service *services.ModpacksService
}

func NewModpacksHandler(service *services.ModpacksService) *ModpacksHandler {
	return &ModpacksHandler{service: service}
}

func (h *ModpacksHandler) GetModpack(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, h.service.Get(id))
}
