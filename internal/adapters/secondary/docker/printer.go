package docker

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func PrintRunningContainers() {
	output, err := GetRunningContainers()
	if err != nil {
		fmt.Printf("Error running docker ps: %v\n", err)
		return
	}

	renderContainerTable(output)
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
