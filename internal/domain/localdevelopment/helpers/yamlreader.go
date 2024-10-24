package helpers

import (
	"fmt"
	"io/ioutil"

	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/vo"
	"gopkg.in/yaml.v3"
)

func LoadProjects(filePath string, projects *vo.Projects) error {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %w", err)
	}

	if err := yaml.Unmarshal(yamlFile, projects); err != nil {
		return fmt.Errorf("error parsing YAML file: %w", err)
	}

	return nil
}
