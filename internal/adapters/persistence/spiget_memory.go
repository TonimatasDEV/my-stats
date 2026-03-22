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

func updateSpiget() {
	for {
		resp, err := http.Get(domain.Spiget)

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var projects []providers.SpigetProject

		if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
			log.Println("Error decoding JSON:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		util.CloseBody(resp.Body)
		processProjectsSpiget(projects)
		time.Sleep(1 * time.Minute)
	}
}

func processProjectsSpiget(projects []providers.SpigetProject) {
	for _, project := range projects {
		dataModrinth[project.Name] = project.Downloads
	}
}
