package docker

import (
	"github.com/mh-daneshvar/dcli/internal/adapters/secondary/cli"
	"github.com/mh-daneshvar/dcli/internal/domain/localdevelopment/vo"
)

func StartProject(project vo.Project) error {
	for _, containerName := range project.Common.Containers {
		if err := cli.RunCommand("docker", []string{"start", containerName}); err != nil {
			return err
		}
	}

	dockerComposeFilePath := project.Common.DockerComposeFilePath
	if dockerComposeFilePath != "" {
		dockerComposeArgs := []string{"-f", dockerComposeFilePath, "up", "-d", "--remove-orphans"}
		if err := cli.RunCommand("docker-compose", dockerComposeArgs); err != nil {
			return err
		}
	}

	return nil
}

func StopProject(project vo.Project) error {
	for _, containerName := range project.Common.Containers {
		if err := cli.RunCommand("docker", []string{"stop", containerName}); err != nil {
			return err
		}
	}

	dockerComposeFilePath := project.Common.DockerComposeFilePath
	if dockerComposeFilePath != "" {
		dockerComposeArgs := []string{"-f", dockerComposeFilePath, "down"}
		if err := cli.RunCommand("docker-compose", dockerComposeArgs); err != nil {
			return err
		}
	}

	for _, service := range project.Services {
		if err := StopService(service); err != nil {
			return err
		}
	}

	return nil
}

func StartService(project vo.Project, service vo.Service) error {
	for _, depSrvLabel := range service.Dependencies {
		service, _, err := project.FindServiceByLabel(depSrvLabel)
		if err != nil {
			return err
		}

		if err := StartService(project, service); err != nil {
			return err
		}
	}

	for _, containerName := range service.Containers {
		if err := cli.RunCommand("docker", []string{"start", containerName}); err != nil {
			return err
		}
	}

	dockerComposeFilePath := service.DockerComposeFilePath
	if dockerComposeFilePath != "" {
		dockerComposeArgs := []string{"-f", dockerComposeFilePath, "up", "-d", "--remove-orphans"}
		if err := cli.RunCommand("docker-compose", dockerComposeArgs); err != nil {
			return err
		}
	}

	return nil
}

func StopService(service vo.Service) error {
	for _, containerName := range service.Containers {
		if err := cli.RunCommand("docker", []string{"stop", containerName}); err != nil {
			return err
		}
	}

	dockerComposeFilePath := service.DockerComposeFilePath
	if dockerComposeFilePath != "" {
		dockerComposeArgs := []string{"-f", dockerComposeFilePath, "down"}
		if err := cli.RunCommand("docker-compose", dockerComposeArgs); err != nil {
			return err
		}
	}

	return nil
}
