package containz

import (
	"testing"
)

func TestRun(t *testing.T) {
	out, err := Run(`
	package main

	import "fmt"

	func main() {
		fmt.Print("Ernesto le Bebecadum")
	}
	`)
	if err != nil {
		t.Fatalf("failed to run %v", err)
	}
	t.Log("Output: ", out)
}
