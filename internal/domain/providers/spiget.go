package providers

import (
	"encoding/json"
	"log"
	"my-stats/internal/domain"
	"my-stats/internal/util"
	"net/http"
)

type SpigetProject struct {
	Name      string `json:"name"`
	Downloads int    `json:"downloads"`
}

func UpdateSpiget(data map[string]int) {
	resp, err := http.Get(domain.Spiget)

	if util.IsNotNil(err) || util.IsNotOk(resp) {
		return
	}

	var projects []SpigetProject

	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	util.CloseBody(resp.Body)
	processProjectsSpiget(data, projects)
}

func processProjectsSpiget(data map[string]int, projects []SpigetProject) {
	for _, project := range projects {
		data[project.Name] = project.Downloads
	}
}
