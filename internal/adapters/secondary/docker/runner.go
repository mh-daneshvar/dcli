package docker

import (
	"fmt"

	"github.com/mh-daneshvar/dcli/internal/adapters/secondary/cli"
	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/vo"
)

func StartProject(project vo.Project) error {
	if err := startContainers(project.Common.Containers); err != nil {
		return fmt.Errorf("failed to start containers for project: %w", err)
	}

	if err := startDockerCompose(project.Common.DockerComposeFilePath); err != nil {
		return fmt.Errorf("failed to start docker-compose for project: %w", err)
	}

	return nil
}

func StopProject(project vo.Project) error {
	if err := stopContainers(project.Common.Containers); err != nil {
		return fmt.Errorf("failed to stop containers for project: %w", err)
	}

	if err := stopDockerCompose(project.Common.DockerComposeFilePath); err != nil {
		return fmt.Errorf("failed to stop docker-compose for project: %w", err)
	}

	for _, service := range project.Services {
		if err := StopService(service); err != nil {
			return fmt.Errorf("failed to stop service %s: %w", service.Label, err)
		}
	}

	return nil
}

func StartService(project vo.Project, service vo.Service) error {
	if err := startDependentServices(project, service); err != nil {
		return fmt.Errorf("failed to start dependencies for service %s: %w", service.Label, err)
	}

	if err := startContainers(service.Containers); err != nil {
		return fmt.Errorf("failed to start containers for service %s: %w", service.Label, err)
	}

	if err := startDockerCompose(service.DockerComposeFilePath); err != nil {
		return fmt.Errorf("failed to start docker-compose for service %s: %w", service.Label, err)
	}

	return nil
}

func StopService(service vo.Service) error {
	if err := stopContainers(service.Containers); err != nil {
		return fmt.Errorf("failed to stop containers for service %s: %w", service.Label, err)
	}

	if err := stopDockerCompose(service.DockerComposeFilePath); err != nil {
		return fmt.Errorf("failed to stop docker-compose for service %s: %w", service.Label, err)
	}

	return nil
}

func startDependentServices(project vo.Project, service vo.Service) error {
	for _, depLabel := range service.Dependencies {
		depService, err := project.FindServiceByLabel(depLabel)
		if err != nil {
			return fmt.Errorf("dependency service not found: %w", err)
		}

		if err := StartService(project, depService); err != nil {
			return fmt.Errorf("failed to start dependent service %s: %w", depLabel, err)
		}
	}
	return nil
}

func startContainers(containers []string) error {
	for _, container := range containers {
		if err := cli.RunCommand("docker", []string{"start", container}); err != nil {
			return fmt.Errorf("failed to start container %s: %w", container, err)
		}
	}
	return nil
}

func stopContainers(containers []string) error {
	for _, container := range containers {
		if err := cli.RunCommand("docker", []string{"stop", container}); err != nil {
			return fmt.Errorf("failed to stop container %s: %w", container, err)
		}
	}
	return nil
}

func startDockerCompose(filePath string) error {
	if filePath == "" {
		return nil
	}
	args := []string{"-f", filePath, "up", "-d", "--remove-orphans"}
	if err := cli.RunCommand("docker-compose", args); err != nil {
		return fmt.Errorf("failed to run docker-compose up: %w", err)
	}
	return nil
}

func stopDockerCompose(filePath string) error {
	if filePath == "" {
		return nil
	}
	args := []string{"-f", filePath, "down"}
	if err := cli.RunCommand("docker-compose", args); err != nil {
		return fmt.Errorf("failed to run docker-compose down: %w", err)
	}
	return nil
}
