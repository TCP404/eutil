package eutil_test

import (
	"os"
	"strings"
	"testing"
)

func TestModulePathUsesLowercaseOwner(t *testing.T) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatalf("ReadFile(go.mod): %v", err)
	}

	firstLine := strings.SplitN(string(data), "\n", 2)[0]
	const want = "module github.com/tcp404/eutil"
	if firstLine != want {
		t.Fatalf("module line = %q, want %q", firstLine, want)
	}
}
