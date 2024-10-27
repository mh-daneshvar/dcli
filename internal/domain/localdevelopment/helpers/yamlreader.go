package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/vo"
)

func LoadProjects(filePath string, projects *vo.Projects) error {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %w", err)
	}

	if err := yaml.Unmarshal(yamlFile, projects); err != nil {
		return fmt.Errorf("error parsing YAML file: %w", err)
	}

	return nil
}
