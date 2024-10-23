package cli

import (
	"github.com/manifoldco/promptui"
	"github.com/mh-daneshvar/dcli/internal/adapters/primary/cli/localdevelopment"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "D-CLI tool",
	Long:  ``,
	Run:   handler,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(localdevelopment.Cmd)
}

func handler(cmd *cobra.Command, args []string) {
	const localDevelopmentOption = "Local Development"

	prompt := promptui.Select{
		Label: "Select An Action",
		Items: []string{localDevelopmentOption},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	switch result {
	case localDevelopmentOption:
		localdevelopment.Cmd.Run(cmd, args)
	}
}
