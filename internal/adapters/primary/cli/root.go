package cli

import (
	"log"

	"github.com/mh-daneshvar/dcli/internal/adapters/primary/cli/localdevelopment"
	"github.com/mh-daneshvar/dcli/internal/common/utils/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "D-CLI tool",
	Long:  "D-CLI is a tool for managing local development environments and other CLI actions.",
	Run:   runRootCmd,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(localdevelopment.Cmd)
}

func runRootCmd(cmd *cobra.Command, args []string) {
	const localDevelopmentOption = "Local Development"

	action, err := ui.RunPrompt(
		"Select an Action",
		[]string{localDevelopmentOption},
	)
	if err != nil {
		log.Fatalf("Failed to run prompt: %v", err)
	}

	switch action {
	case localDevelopmentOption:
		localdevelopment.Cmd.Run(cmd, args)
	default:
		log.Printf("Unknown option selected: %s", action)
	}
}
