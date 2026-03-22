package providers

import (
	"encoding/json"
	"log"
	"my-stats/internal/domain"
	"my-stats/internal/util"
	"net/http"
)

type HangarProjects struct {
	Result []HangarProject `json:"result"`
}

type HangarProject struct {
	Name  string `json:"name"`
	Stats struct {
		Downloads int `json:"downloads"`
	} `json:"stats"`
}

func UpdateHangar(data map[string]int) {
	resp, err := http.Get(domain.Hangar)

	if util.IsNotNil(err) || util.IsNotOk(resp) {
		return
	}

	var projects HangarProjects

	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	util.CloseBody(resp.Body)
	processProjectsHangar(data, projects.Result)
}

func processProjectsHangar(data map[string]int, projects []HangarProject) {
	for _, project := range projects {
		data[project.Name] = project.Stats.Downloads
	}
}
