package cli

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/schollz/progressbar/v3"
)

func RunCommand(command string, args []string) error {
	bar := createProgressBar("Running commands...")

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start commands: %w", err)
	}

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	return monitorCommandProgress(done, bar)
}

func createProgressBar(description string) *progressbar.ProgressBar {
	return progressbar.NewOptions(-1,
		progressbar.OptionSetDescription(description),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionFullWidth(),
		progressbar.OptionClearOnFinish(),
	)
}

func monitorCommandProgress(done chan error, bar *progressbar.ProgressBar) error {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case err := <-done:
			_ = bar.Finish()
			if err != nil {
				return fmt.Errorf("commands execution failed: %w", err)
			}
			return nil
		case <-ticker.C:
			_ = bar.Add(1)
		}
	}
}
