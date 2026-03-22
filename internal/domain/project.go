package domain

type Project struct {
	Name      string `json:"name"`
	Downloads int    `json:"downloads"`
}

func SendArray(projects []Project) {
	// TODO
}
