//go:build !inmemory
// +build !inmemory

package pomodoro_test

import (
	"os"
	"small_console_applications_go/interactiveTools/pomo/pomodoro"
	"small_console_applications_go/interactiveTools/pomo/pomodoro/repository"
	"testing"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	tf, err := os.CreateTemp("", "pomo")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()
	dbRepo, err := repository.NewSQLite3Repo(tf.Name())
	if err != nil {
		t.Fatal(err)
	}
	return dbRepo, func() {
		os.Remove(tf.Name())
	}
}
