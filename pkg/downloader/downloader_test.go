package downloader

import (
	"os"
	"strings"
	"testing"
)

func TestDownload_NotFound(t *testing.T) {
	orig := os.Getenv("PATH")
	t.Cleanup(func() { os.Setenv("PATH", orig) })
	os.Setenv("PATH", "")

	err := Download("https://example.com/song")
	if err == nil {
		t.Fatal("expected an error when yt-dlp is not in PATH, got nil")
	}
	if !strings.Contains(err.Error(), "yt-dlp not found") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestDownload_Success(t *testing.T) {
	dir := t.TempDir()

	fakeBin := dir + "/yt-dlp"
	script := "#!/bin/sh\nexit 0\n"
	if err := os.WriteFile(fakeBin, []byte(script), 0o755); err != nil {
		t.Fatalf("could not create fake yt-dlp: %v", err)
	}

	orig := os.Getenv("PATH")
	t.Cleanup(func() { os.Setenv("PATH", orig) })
	os.Setenv("PATH", dir+":"+orig)

	if err := Download("https://www.youtube.com/watch?v=9sJUDx7iEJw"); err != nil {
		t.Fatalf("expected no error with a successful yt-dlp, got: %v", err)
	}
}

func TestDownload_Failure(t *testing.T) {
	dir := t.TempDir()

	fakeBin := dir + "/yt-dlp"
	script := "#!/bin/sh\nexit 1\n"
	if err := os.WriteFile(fakeBin, []byte(script), 0o755); err != nil {
		t.Fatalf("could not create fake yt-dlp: %v", err)
	}

	orig := os.Getenv("PATH")
	t.Cleanup(func() { os.Setenv("PATH", orig) })
	os.Setenv("PATH", dir+":"+orig)

	err := Download("https://example.com/bad-url")
	if err == nil {
		t.Fatal("expected an error when yt-dlp exits non-zero, got nil")
	}
	if !strings.Contains(err.Error(), "yt-dlp failed") {
		t.Errorf("unexpected error message: %v", err)
	}
}
