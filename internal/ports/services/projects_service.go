package services

import (
	"my-stats/internal/domain"
	"my-stats/internal/ports/repositories"
)

type ProjectsService struct {
	projectsRepo repositories.ProjectsRepository
}

func NewProjectsService(projectsRepo repositories.ProjectsRepository) *ProjectsService {
	return &ProjectsService{projectsRepo: projectsRepo}
}

func (service *ProjectsService) Get() []domain.Project {
	var projects []domain.Project

	for name, downloads := range service.projectsRepo.Get() {
		projects = append(projects, domain.Project{
			Name:      name,
			Downloads: downloads,
		})
	}

	return projects
}
