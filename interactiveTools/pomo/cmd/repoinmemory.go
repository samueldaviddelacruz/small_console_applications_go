//go:build inmemory
// +build inmemory

package cmd

import (
	"small_console_applications_go/interactiveTools/pomo/pomodoro"
	"small_console_applications_go/interactiveTools/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
