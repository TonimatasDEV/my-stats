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

var dataCurseForge = make(map[string]int)

func updateCurseForge() {
	for {
		resp, err := http.Get(domain.CFWidget)

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var author providers.CurseForgeAuthor

		if err := json.NewDecoder(resp.Body).Decode(&author); err != nil {
			log.Println("Error decoding JSON:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		util.CloseBody(resp.Body)
		processProjectsCurseForge(author.Projects)
		time.Sleep(1 * time.Minute)
	}
}

func processProjectsCurseForge(projects []providers.CurseForgeProject) {
	for _, providerProject := range projects {
		resp, err := http.Get(domain.CFWidgetProject + strconv.Itoa(providerProject.ID))

		if util.IsNotNil(err) || util.IsNotOk(resp) {
			continue
		}

		var projectInfo providers.CurseForgeProjectInfo

		if err := json.NewDecoder(resp.Body).Decode(&projectInfo); err != nil {
			log.Println("Error decoding JSON:", err)
			continue
		}

		dataCurseForge[providerProject.Name] = projectInfo.Downloads.Total
		util.CloseBody(resp.Body)
	}
}
