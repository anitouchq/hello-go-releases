package greeting

import "testing"

func TestMessage(t *testing.T) {
	if got := Message(); got == "" {
		t.Fatal("Message() must not be empty")
	}
}

func TestBuildInfo(t *testing.T) {
	got := BuildInfo("v1.0.0", "linux", "amd64")
	want := "Version: v1.0.0\nOS: linux\nArch: amd64"
	if got != want {
		t.Fatalf("BuildInfo() = %q, want %q", got, want)
	}
}
