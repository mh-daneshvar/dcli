package vo

import (
	"fmt"
)

type Project struct {
	Services []Service `yaml:"services"`
	Common   Common    `yaml:"common"`
}

func (p Project) GetServiceLabels() []string {
	services := p.Services

	labels := make([]string, 0, len(services))
	for _, service := range services {
		labels = append(labels, service.Label)
	}

	return labels
}

func (p Project) FindServiceByLabel(label string) (Service, bool, error) {
	for _, service := range p.Services {
		if service.Label == label {
			return service, false, nil
		}
	}
	return Service{}, false, fmt.Errorf("selected service not found: %s", label)
}
