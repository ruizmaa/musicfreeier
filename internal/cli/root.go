package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "musicfreeier <url>",
	Short:        "Download music from any URL",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true, // don't print usage on runtime errors
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		return download(url)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func download(url string) error {
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
