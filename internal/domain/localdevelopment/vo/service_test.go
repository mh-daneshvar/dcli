package vo

import (
	"testing"
)

func TestServiceStruct(t *testing.T) {
	service := Service{
		Label:                 "test-service",
		Containers:            []string{"container1", "container2"},
		DockerComposeFilePath: "/path/to/docker-compose.yml",
		Dependencies:          []string{"dependency1"},
	}

	if service.Label != "test-service" {
		t.Errorf("unexpected label: %s", service.Label)
	}
	if len(service.Containers) != 2 {
		t.Errorf("expected 2 containers, got %d", len(service.Containers))
	}
	if len(service.Dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(service.Dependencies))
	}
}
