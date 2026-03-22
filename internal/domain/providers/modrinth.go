package providers

import (
	"encoding/json"
	"log"
	"my-stats/internal/domain"
	"my-stats/internal/util"
	"net/http"
)

type ModrinthProject struct {
	Title     string `json:"title"`
	Downloads int    `json:"downloads"`
}

func UpdateModrinth(data map[string]int) {
	resp, err := http.Get(domain.Modrinth)

	if util.IsNotNil(err) || util.IsNotOk(resp) {
		return
	}

	var projects []ModrinthProject

	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	util.CloseBody(resp.Body)
	processProjectsModrinth(data, projects)
}

func processProjectsModrinth(data map[string]int, projects []ModrinthProject) {
	for _, project := range projects {
		data[project.Title] = project.Downloads
	}
}
