package persistence

import (
	"encoding/json"
	"fmt"
	"my-stats/internal/domain"
	"my-stats/internal/domain/providers"
	"my-stats/internal/util"
	"net/http"
	"strconv"
	"time"
)

var data = make(map[string]int)

func Update() {
	for {
		resp, err := http.Get(domain.CFWidget)

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var author providers.ProviderAuthor

		if err := json.NewDecoder(resp.Body).Decode(&author); err != nil {
			fmt.Println("Error decoding JSON:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		util.CloseBody(resp.Body)
		processProjects(author.Projects)
		time.Sleep(1 * time.Minute)
	}
}

func processProjects(projects []providers.ProviderProject) {
	for _, providerProject := range projects {
		resp, err := http.Get(domain.CFWidgetProject + strconv.Itoa(providerProject.ID))

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var projectInfo providers.ProviderProjectInfo

		if err := json.NewDecoder(resp.Body).Decode(&projectInfo); err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}

		data[providerProject.Name] = projectInfo.Downloads.Total
		util.CloseBody(resp.Body)
	}
}
