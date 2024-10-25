package vo

import (
	"testing"
)

func TestGetProjectNames(t *testing.T) {
	projects := Projects{
		Projects: map[string]Project{
			"project1": {},
			"project2": {},
		},
	}

	projectNames := projects.GetProjectNames()
	expectedLength := 2
	if len(projectNames) != expectedLength {
		t.Errorf("expected %d project names, got %d", expectedLength, len(projectNames))
	}
}

func TestFindProjectByName(t *testing.T) {
	projects := Projects{
		Projects: map[string]Project{
			"project1": {},
			"project2": {},
		},
	}

	_, err := projects.FindProjectByName("project1")
	if err != nil {
		t.Errorf("expected project1 to be found, but got error: %v", err)
	}

	_, err = projects.FindProjectByName("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent project, but got none")
	}
}
