package vo

import (
	"testing"
)

func TestCommonStruct(t *testing.T) {
	common := Common{
		Containers:            []string{"container1", "container2"},
		DockerComposeFilePath: "/path/to/docker-compose.yml",
	}

	if len(common.Containers) != 2 {
		t.Errorf("expected 2 containers, got %d", len(common.Containers))
	}
	if common.DockerComposeFilePath != "/path/to/docker-compose.yml" {
		t.Errorf("unexpected DockerComposeFilePath: %s", common.DockerComposeFilePath)
	}
}
