package helpers

import (
	"testing"
)

func TestBuildVersionReturnsDevWhenEmpty(t *testing.T) {
	Version = ""
	result := BuildVersion()
	expected := "dev"

	if result != expected {
		t.Fatalf("expected %s got %s", expected, result)
	}
}

func TestBuildVersionReturnsVersionWhenAvailable(t *testing.T) {
	expected := "version"
	Version = expected
	result := BuildVersion()

	if result != expected {
		t.Fatalf("expected %s got %s", expected, result)
	}
}
