// Package downloader provides functionality to download audio from URLs
// using yt-dlp. It can be used as a library by other Go modules.
package downloader

import (
	"fmt"
	"os"
	"os/exec"
)

// Download downloads audio from the given URL as an MP3 file using yt-dlp.
// The output file is named after the title of the media and placed in the
// current working directory.
//
// It requires yt-dlp to be installed and available in PATH.
func Download(url string) error {
	fmt.Printf("⬇  Downloading: %s\n", url)

	ytdlp, err := exec.LookPath("yt-dlp")
	if err != nil {
		return fmt.Errorf("yt-dlp not found in PATH: %w", err)
	}

	cmd := exec.Command(ytdlp,
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "0", // best quality
		"--output", "%(title)s.%(ext)s",
		url,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp failed: %w", err)
	}

	return nil
}
