package http

import (
	"context"
	"log"
	"testing"
)

func f1(ctx context.Context, input *Input, result *Result) {
	log.Println("f1", ctx, input, result)
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
			log.Println("api-test", ctx, input, result)
		}))
		task.Run()
	})
}
