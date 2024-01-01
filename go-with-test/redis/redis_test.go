package main

import "testing"

func TestRedis(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		t.Helper()
		example()
	})

	t.Run("bad", func(t *testing.T) {
		t.Helper()
		lockTest()
	})
}
