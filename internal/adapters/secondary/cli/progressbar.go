package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

func RunCommand(command string, args []string) error {
	bar := createProgressBar("Running commands...")
	cmd := prepareCommand(command, args)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	done := make(chan error, 1)
	go monitorCommandExecution(cmd, done)

	return handleCommandProgress(done, bar)
}

func RunCliCommand(inputCommand string) error {
	printHeader(inputCommand)
	parts := parseCommand(inputCommand)
	return RunCommand(parts[0], parts[1:])
}

func prepareCommand(command string, args []string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func monitorCommandExecution(cmd *exec.Cmd, done chan<- error) {
	done <- cmd.Wait()
}

func handleCommandProgress(done chan error, bar *progressbar.ProgressBar) error {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case err := <-done:
			if finishErr := bar.Finish(); finishErr != nil {
				return fmt.Errorf("error finishing progress bar: %w", finishErr)
			}
			if err != nil {
				return fmt.Errorf("command execution failed: %w", err)
			}
			return nil
		case <-ticker.C:
			if addErr := bar.Add(1); addErr != nil {
				return fmt.Errorf("error updating progress bar: %w", addErr)
			}
		}
	}
}

func parseCommand(input string) []string {
	return strings.Fields(input)
}
