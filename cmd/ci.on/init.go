package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func (service mainService) RunInit() error {
	service.logService.Info("Initializing CI.On for GitLab...")

	// verifica se gitlab-ci-local já está instalado
	_, err := exec.LookPath("gitlab-ci-local")
	if err == nil {
		service.logService.Info("gitlab-ci-local is already installed ✅")
		return nil
	}

	// instala via npm global
	service.logService.Info("Installing gitlab-ci-local...")
	var installCmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		// comando para macOS
		installCmd = exec.Command("brew", "install", "gitlab-ci-local")
	case "windows":
		// comando para windows
		installCmd = exec.Command("cmd", "/C", "npm", "install", "-g", "gitlab-ci-local")
	default:
		// comando para linux
		installCmd = exec.Command("npm", "install", "-g", "gitlab-ci-local")
	}
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		service.logService.Error(500, "install_gitlab_ci_local", err.Error())
		return fmt.Errorf("failed to install gitlab-ci-local: %w", err)
	}

	service.logService.Info("gitlab-ci-local installed successfully ✅")
	return nil

}
