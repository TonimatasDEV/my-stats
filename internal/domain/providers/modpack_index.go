package providers

import (
	"encoding/json"
	"log"
	"my-stats/internal/domain"
	"my-stats/internal/util"
	"net/http"
	"strings"
)

type ModpackIndexInfo struct {
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
}

var hardcoded = []string{"39585", "77177"}

func UpdateModpackIndex(data map[string]domain.Modpack) {
	resp, err := http.Get(domain.ModpackIndexModpacks)

	if util.IsNotNil(err) || util.IsNotOk(resp) {
		return
	}

	var info ModpackIndexInfo

	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	util.CloseBody(resp.Body)
	processProjectsModpackIndex(data, hardcoded, info)
}

func processProjectsModpackIndex(data map[string]domain.Modpack, projects []string, info ModpackIndexInfo) {
	for _, project := range projects {
		url := strings.ReplaceAll(domain.ModpackIndexModData, "{id}", project)
		resp, err := http.Get(url)

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			return
		}

		var modDataInfo ModpackIndexInfo

		if err := json.NewDecoder(resp.Body).Decode(&modDataInfo); err != nil {
			log.Println("Error decoding JSON:", err)
			return
		}

		data[project] = domain.Modpack{
			Total: info.Meta.Total,
			With:  modDataInfo.Meta.Total,
		}
	}
}
