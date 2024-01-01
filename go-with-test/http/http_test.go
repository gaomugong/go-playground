package http

import (
	"context"
	"testing"
)

func f1(ctx context.Context, input *Input, result *Result) {
}

func TestTask(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		t.Helper()

		task := &Task{
			name:    "cleaner",
			runners: make(map[string]Runner),
		}

		task.AddRunner("clean-task", RunnerFunc(f1))
		task.AddRunner("api-test", RunnerFunc(func(ctx context.Context, input *Input, result *Result) {
		}))
		task.Run()
	})
}
