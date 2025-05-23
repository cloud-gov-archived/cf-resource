package out

import resource "github.com/cloud-gov/cf-resource"

type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	ManifestPath         string                 `json:"manifest"`
	Path                 string                 `json:"path"`
	CurrentAppName       string                 `json:"current_app_name"`
	Vars                 map[string]interface{} `json:"vars"`
	VarsFiles            []string               `json:"vars_files"`
	EnvironmentVariables map[string]string      `json:"environment_variables"`
	DockerUsername       string                 `json:"docker_username"`
	DockerPassword       string                 `json:"docker_password"`
	ShowAppLog           bool                   `json:"show_app_log"`
	NoStart              bool                   `json:"no_start"`
	Task                 bool                   `json:"task"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
