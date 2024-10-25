package vo

import "fmt"

type Projects struct {
	Projects map[string]Project `yaml:"projects"`
}

func (projects Projects) GetProjectNames() []string {
	projectNames := make([]string, 0, len(projects.Projects))
	for projectName := range projects.Projects {
		projectNames = append(projectNames, projectName)
	}
	return projectNames
}

func (projects Projects) FindProjectByName(selectedProjectName string) (Project, error) {
	project, exists := projects.Projects[selectedProjectName]
	if !exists {
		return Project{}, fmt.Errorf("selected project not found: %s", selectedProjectName)
	}
	return project, nil
}
