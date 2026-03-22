package persistence

import (
	"encoding/json"
	"log"
	"my-stats/internal/domain"
	"my-stats/internal/domain/providers"
	"my-stats/internal/util"
	"net/http"
	"strconv"
	"time"
)

var dataCFWidget = make(map[string]int)

func updateCFWidget() {
	for {
		resp, err := http.Get(domain.CFWidget)

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var author providers.CFWidgetAuthor

		if err := json.NewDecoder(resp.Body).Decode(&author); err != nil {
			log.Println("Error decoding JSON:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		util.CloseBody(resp.Body)
		processProjectsCFWidget(author.Projects)
		time.Sleep(1 * time.Minute)
	}
}

func processProjectsCFWidget(projects []providers.CFWidgetProject) {
	for _, providerProject := range projects {
		resp, err := http.Get(domain.CFWidgetProject + strconv.Itoa(providerProject.ID))

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var projectInfo providers.CFWidgetProjectInfo

		if err := json.NewDecoder(resp.Body).Decode(&projectInfo); err != nil {
			log.Println("Error decoding JSON:", err)
			continue
		}

		dataCFWidget[providerProject.Name] = projectInfo.Downloads.Total
		util.CloseBody(resp.Body)
	}
}
