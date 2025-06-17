package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-repo",
	Short: "Analyze GitHub repositories and fetch detailed information",
	Long: `go-repo is a CLI tool to analyze GitHub repositories and fetch comprehensive information
including stars, forks, issues, pull requests, and repository metadata.

The tool supports multiple output formats (table, JSON) and can save results to files
for further analysis or reporting.`,
	Example: `  # Basic repository analysis
  go-repo analyze golang/go
  go-repo analyze https://github.com/microsoft/vscode

  # With pull requests
  go-repo analyze golang/go --prs
  go-repo analyze golang/go --prs 15

  # Different output formats
  go-repo analyze golang/go --format json
  go-repo analyze golang/go --format table

  # Download results to files
  go-repo analyze golang/go --download json
  go-repo analyze golang/go --download pdf --prs 10

  # Using GitHub token for higher rate limits
  go-repo analyze golang/go --token ghp_xxxxxxxxxxxx
  go-repo analyze golang/go -t ghp_xxxxxxxxxxxx

  # Complete analysis with all options
  go-repo analyze microsoft/vscode --prs 20 --format json --download json --token ghp_xxxxxxxxxxxx`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "",
		`GitHub personal access token for authenticated API requests.
Can also be set via GITHUB_TOKEN environment variable.
Without a token, you're limited to 60 requests per hour.
With a token, you get 5000 requests per hour.

Get your token at: https://github.com/settings/tokens`)
}

var token string
