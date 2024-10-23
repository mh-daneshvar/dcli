package helpers

import (
	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/vo"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func LoadProjects(filePath string, projects *vo.Projects) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
	}

	err = yaml.Unmarshal(yamlFile, projects)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s", err)
	}
}
