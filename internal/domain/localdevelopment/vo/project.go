package vo

import (
	"fmt"
)

type Project struct {
	Services []Service `yaml:"services"`
	Common   Common    `yaml:"common"`
}

func (p *Project) GetServiceLabels() []string {
	labels := make([]string, 0, len(p.Services))
	for _, service := range p.Services {
		labels = append(labels, service.Label)
	}
	return labels
}

func (p *Project) FindServiceByLabel(label string) (Service, error) {
	for _, service := range p.Services {
		if service.Label == label {
			return service, nil
		}
	}
	return Service{}, fmt.Errorf("service with label %s not found", label)
}
