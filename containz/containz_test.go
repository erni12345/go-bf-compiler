package containz

import (
	"testing"
)

func TestRun(t *testing.T) {
	_, err := Run("example")
	if err != nil {
		t.Fatalf("failed to run %v", err)
	}
}
