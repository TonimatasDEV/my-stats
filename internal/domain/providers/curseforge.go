package providers

type ProviderProject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProviderAuthor struct {
	Projects []ProviderProject `json:"projects"`
}

type ProviderProjectInfo struct {
	Downloads struct {
		Total int `json:"total"`
	} `json:"downloads"`
}
