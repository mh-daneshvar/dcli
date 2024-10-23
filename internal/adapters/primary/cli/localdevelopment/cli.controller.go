package localdevelopment

import (
	"fmt"
	"github.com/mh-daneshvar/dcli/internal/common/utils/ui"
	"github.com/spf13/cobra"
	"log"

	"github.com/mh-daneshvar/dcli/internal/adapters/secondary/docker"
	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/helpers"
	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/vo"
)

var (
	projects vo.Projects
	Cmd      = createCobraCommand()
)

func createCobraCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "",
		Short: "Local Development",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := handler(); err != nil {
				log.Fatalf("Error in handler: %v", err)
			}
		},
	}
}

func init() {
	helpers.LoadProjects("configs/local-development.yaml", &projects)
}

func handler() error {
	project, err := selectProject(projects)
	if err != nil {
		return fmt.Errorf("error selecting project: %w", err)
	}

	service, stopAllSelected, err := selectService(project, "Stop All Services")
	if err != nil {
		return fmt.Errorf("error selecting service: %w", err)
	}

	if stopAllSelected {
		docker.StopProject(project)
	} else {
		docker.StartProject(project)
		docker.StartService(project, service)
	}

	docker.PrintRunningContainers()
	return nil
}

func selectProject(projects vo.Projects) (vo.Project, error) {
	projectNames := projects.GetProjectNames()

	selectedProjectName, err := ui.RunPrompt("Select the Project", projectNames)
	if err != nil {
		return vo.Project{}, fmt.Errorf("prompt failed: %w", err)
	}

	return projects.FindProjectByName(selectedProjectName)
}

func selectService(project vo.Project, stopAllLabel string) (vo.Service, bool, error) {
	serviceLabels := project.GetServiceLabels()
	serviceNames := append(serviceLabels, stopAllLabel)

	selectedServiceLabel, err := ui.RunPrompt("Select the Service", serviceNames)
	if err != nil {
		return vo.Service{}, false, fmt.Errorf("prompt failed: %w", err)
	}

	if selectedServiceLabel == stopAllLabel {
		return vo.Service{}, true, nil
	}

	return project.FindServiceByLabel(selectedServiceLabel)
}
