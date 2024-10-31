package localdevelopment

import (
	"fmt"
	"github.com/mh-daneshvar/dcli/internal/adapters/secondary/cli"
	"github.com/spf13/cobra"
	"log"

	"github.com/mh-daneshvar/dcli/internal/adapters/secondary/docker"
	"github.com/mh-daneshvar/dcli/internal/common/utils/ui"
	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/helpers"
	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/vo"
)

var (
	projects vo.Projects
	Cmd      *cobra.Command
)

func init() {
	Cmd = createCobraCommand()

	if err := helpers.LoadProjects("configs/local-development.yaml", &projects); err != nil {
		log.Fatalf("Failed to load projects: %v", err)
	}
}

func createCobraCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "",
		Short: "Manage local development environments",
		Long:  "This command helps in managing local development environments including starting and stopping services.",
		Run:   runHandler,
	}
}

func runHandler(cmd *cobra.Command, args []string) {
	if err := handler(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func handler() error {
	project, err := selectProject(projects)
	if err != nil {
		return fmt.Errorf("failed to select project: %w", err)
	}

	service, stopAllSelected, err := selectService(project, "Stop All Services")
	if err != nil {
		return fmt.Errorf("failed to select service: %w", err)
	}

	if stopAllSelected {
		if err := docker.StopProject(project); err != nil {
			return fmt.Errorf("failed to stop project: %w", err)
		}
	} else {
		if err := docker.StartProject(project); err != nil {
			return fmt.Errorf("failed to start project: %w", err)
		}
		if err := docker.StartService(project, service); err != nil {
			return fmt.Errorf("failed to start service: %w", err)
		}
		for _, command := range service.Commands {
			if err := cli.RunCliCommand(command); err != nil {
				return fmt.Errorf("failed to run command: %w", err)
			}
		}
	}

	docker.PrintRunningContainers()
	return nil
}

func selectProject(projects vo.Projects) (vo.Project, error) {
	projectNames := projects.GetProjectNames()

	selectedProjectName, err := ui.RunPrompt("Select the Project", projectNames)
	if err != nil {
		return vo.Project{}, fmt.Errorf("failed to run project selection prompt: %w", err)
	}

	project, err := projects.FindProjectByName(selectedProjectName)
	if err != nil {
		return vo.Project{}, fmt.Errorf("failed to find project: %w", err)
	}
	return project, nil
}

func selectService(project vo.Project, stopAllLabel string) (vo.Service, bool, error) {
	serviceLabels := project.GetServiceLabels()
	serviceOptions := append(serviceLabels, stopAllLabel)

	selectedServiceLabel, err := ui.RunPrompt("Select the Service", serviceOptions)
	if err != nil {
		return vo.Service{}, false, fmt.Errorf("failed to run service selection prompt: %w", err)
	}

	if selectedServiceLabel == stopAllLabel {
		return vo.Service{}, true, nil
	}

	service, err := project.FindServiceByLabel(selectedServiceLabel)
	if err != nil {
		return vo.Service{}, false, fmt.Errorf("failed to find service: %w", err)
	}
	return service, false, nil
}
