package providers

import (
	"encoding/json"
	"log"
	"my-stats/internal/domain"
	"my-stats/internal/util"
	"net/http"
	"strconv"
)

type CFWidgetProject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CFWidgetAuthor struct {
	Projects []CFWidgetProject `json:"projects"`
}

type CFWidgetProjectInfo struct {
	Downloads struct {
		Total int `json:"total"`
	} `json:"downloads"`
}

func UpdateCFWidget(data map[string]int) {
	resp, err := http.Get(domain.CFWidget)

	if util.IsNotNil(err) || util.IsNotOk(resp) {
		return
	}

	var author CFWidgetAuthor

	if err := json.NewDecoder(resp.Body).Decode(&author); err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	util.CloseBody(resp.Body)
	processProjectsCFWidget(data, author.Projects)
}

func processProjectsCFWidget(data map[string]int, projects []CFWidgetProject) {
	for _, providerProject := range projects {
		resp, err := http.Get(domain.CFWidgetProject + strconv.Itoa(providerProject.ID))

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var projectInfo CFWidgetProjectInfo

		if err := json.NewDecoder(resp.Body).Decode(&projectInfo); err != nil {
			log.Println("Error decoding JSON:", err)
			continue
		}

		data[providerProject.Name] = projectInfo.Downloads.Total
		util.CloseBody(resp.Body)
	}
}
