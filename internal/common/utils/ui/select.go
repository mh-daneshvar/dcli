package ui

import "github.com/manifoldco/promptui"

func RunPrompt(label string, options []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
