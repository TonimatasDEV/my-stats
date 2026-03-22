package persistence

import (
	"encoding/json"
	"fmt"
	"my-stats/internal/domain"
	"my-stats/internal/domain/providers"
	"my-stats/internal/util"
	"net/http"
	"time"
)

var dataModrinth = make(map[string]int)

func updateModrinth() {
	for {
		resp, err := http.Get(domain.Modrinth)

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var projects []providers.ModrinthProject

		if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
			fmt.Println("Error decoding JSON:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		util.CloseBody(resp.Body)
		processProjectsModrinth(projects)
		time.Sleep(1 * time.Minute)
	}
}

func processProjectsModrinth(projects []providers.ModrinthProject) {
	for _, project := range projects {
		dataModrinth[project.Title] = project.Downloads
	}
}
