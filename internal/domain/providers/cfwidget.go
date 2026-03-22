package providers

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
