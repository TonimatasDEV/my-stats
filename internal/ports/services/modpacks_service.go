package services

import (
	"my-stats/internal/domain"
	"my-stats/internal/ports/repositories"
)

type ModpacksService struct {
	projectsRepo repositories.ModpacksRepository
}

func NewModpacksService(projectsRepo repositories.ModpacksRepository) *ModpacksService {
	return &ModpacksService{projectsRepo: projectsRepo}
}

func (service *ModpacksService) Get(id string) domain.Modpack {
	return service.projectsRepo.Get()[id]
}
