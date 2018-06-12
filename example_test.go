package trylock_test

import (
	"testing"

	"github.com/chappjc/trylock"
)

func TestExample(t *testing.T) {
	var mu trylock.Mutex
	t.Log(mu.TryLock())
	t.Log(mu.TryLock())
	mu.Unlock()
	t.Log(mu.TryLock())
	// Output:
	// true
	// false
	// true
}
