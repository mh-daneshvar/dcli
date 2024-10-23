package vo

type Common struct {
	Containers            []string `yaml:"containers"`
	DockerComposeFilePath string   `yaml:"docker_compose_file_path"`
}
