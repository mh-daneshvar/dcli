package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func PrintRunningContainers() {
	output, err := getDockerContainers()
	if err != nil {
		fmt.Printf("Error running docker ps: %v\n", err)
		return
	}

	renderContainerTable(output)
}

func getDockerContainers() (string, error) {
	cmd := exec.Command("docker", "ps", "--format", "{{.Names}}\\t{{.Image}}")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running docker ps: %w", err)
	}
	return string(output), nil
}

func renderContainerTable(output string) {
	lines := strings.Split(output, "\n")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"CONTAINER NAME", "IMAGE"})

	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Split(line, "\t")
		table.Append(fields)
	}

	table.Render()
}
