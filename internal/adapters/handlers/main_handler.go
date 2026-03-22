package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainMessage struct {
	Msg          string `json:"msg"`
	Wiki         string `json:"wiki"`
	License      string `json:"license"`
	Repository   string `json:"repository"`
	IssueTracker string `json:"issueTracker"`
}

func HandleMain(c *gin.Context) {
	mainMessage := MainMessage{
		Msg:          "Hello World! This is an API to get my projects stats",
		Wiki:         "https://github.com/TonimatasDEV/my-stats",
		License:      "https://github.com/TonimatasDEV/my-stats/blob/master/LICENSE",
		Repository:   "https://github.com/TonimatasDEV/my-stats",
		IssueTracker: "https://github.com/TonimatasDEV/my-stats/issues",
	}

	c.JSON(http.StatusOK, mainMessage)
}
