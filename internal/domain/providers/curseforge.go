package providers

type CurseForgeProject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CurseForgeAuthor struct {
	Projects []CurseForgeProject `json:"projects"`
}

type CurseForgeProjectInfo struct {
	Downloads struct {
		Total int `json:"total"`
	} `json:"downloads"`
}
