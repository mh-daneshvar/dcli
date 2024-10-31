package cli

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/schollz/progressbar/v3"
	"os"
)

func printHeader(header string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#D845C0")).
		PaddingLeft(1).
		PaddingRight(1).
		MarginTop(1).
		MarginBottom(1)

	fmt.Println(style.Render(header))
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
