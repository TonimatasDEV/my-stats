package providers

type HangarProjects struct {
	Result []HangarProject `json:"result"`
}

type HangarProject struct {
	Name  string `json:"name"`
	Stats struct {
		Downloads int `json:"downloads"`
	} `json:"stats"`
}
