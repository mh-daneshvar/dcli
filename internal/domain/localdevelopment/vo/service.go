package vo

type Service struct {
	Label                 string   `yaml:"label"`
	Containers            []string `yaml:"containers"`
	DockerComposeFilePath string   `yaml:"docker_compose_file_path"`
	Dependencies          []string `yaml:"dependencies"`
}
