//go:build !inmemory
// +build !inmemory

package cmd

import (
	"small_console_applications_go/interactiveTools/pomo/pomodoro"
	"small_console_applications_go/interactiveTools/pomo/pomodoro/repository"

	"github.com/spf13/viper"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQlite3Repo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}
	return repo, nil
}
