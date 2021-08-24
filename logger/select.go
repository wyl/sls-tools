package logger

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func YesNo(label string) bool {
	prompt := promptui.Select{
		Label: fmt.Sprintf("%v [Yes/No]", label),
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		MyLogger.ErrorLog.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}
