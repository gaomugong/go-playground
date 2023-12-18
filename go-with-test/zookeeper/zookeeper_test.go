package zookeeper

import "testing"

func TestZk(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		t.Helper()
		main()
	})

	t.Run("basic test", func(t *testing.T) {
		connectBasic()
	})

	t.Run("advanced test", func(t *testing.T) {
		connectAdvance()
	})
}
