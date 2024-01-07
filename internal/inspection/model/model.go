package model

type App struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Pod     Pod    `json:"pod"`
}

type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	IP        string `json:"ip"`
}
