package main

import "testing"

func TestRun(t *testing.T) {
	_, err := run()
	if err != nil {
		t.Error("failed run()")
	}
}

// how to run tests :

// go test
// go test -v
// go test -cover for see coverage
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out for see coverage in browser
