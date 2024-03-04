package pomodoro_test

import (
	"small_console_applications_go/interactiveTools/pomo/pomodoro"
	"small_console_applications_go/interactiveTools/pomo/pomodoro/repository"
	"testing"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
