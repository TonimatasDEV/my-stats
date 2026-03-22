package persistence

import (
	"encoding/json"
	"log"
	"my-stats/internal/domain"
	"my-stats/internal/domain/providers"
	"my-stats/internal/util"
	"net/http"
	"time"
)

func updateHangar() {
	for {
		resp, err := http.Get(domain.Hangar)

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var projects providers.HangarProjects

		if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
			log.Println("Error decoding JSON:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		util.CloseBody(resp.Body)
		processProjectsHangar(projects.Result)
		time.Sleep(1 * time.Minute)
	}
}

func processProjectsHangar(projects []providers.HangarProject) {
	for _, project := range projects {
		dataModrinth[project.Name] = project.Stats.Downloads
	}
}
