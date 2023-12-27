package controller

import (
	"os/exec"
)

func GenerateModel() string {
	_, err := exec.Command("touch", "models.go").Output()
	if err != nil {
		return err.Error()
	}
	return "models.go created successfully"
}
