package cli

import (
	"fmt"
	"os"

	"github.com/ruizmaa/musicfreeier/pkg/downloader"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "musicfreeier <url>",
	Short:        "Download music from any URL",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true, // don't print usage on runtime errors
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		return downloader.Download(url)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
