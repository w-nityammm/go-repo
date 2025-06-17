package cmd

import (
	"log"

	"go-repo/internal/analyzer"
	"go-repo/internal/output"

	"github.com/spf13/cobra"
)

var (
	format   string
	prs      int
	download string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze [repository]",
	Short: "Analyze a specific GitHub repository",
	Long: `Analyze a GitHub repository and display comprehensive information including:
- Repository metadata (name, description, language)
- Statistics (stars, forks, open issues)
- Timestamps (created, last updated)
- Recent pull requests (optional)

The repository can be specified in two formats:
  1. Short format: owner/repo (e.g., golang/go)
  2. Full URL: https://github.com/owner/repo

Results can be displayed in multiple formats.`,
	Args: cobra.ExactArgs(1),
	Run:  runAnalyze,
	Example: `  # Basic analysis (table format, no PRs)
  go-repo analyze golang/go
  go-repo analyze https://github.com/microsoft/vscode

  # Show pull requests (default 5 when --prs used without number)
  go-repo analyze golang/go --prs
  go-repo analyze golang/go -p

  # Show specific number of pull requests
  go-repo analyze golang/go --prs 15
  go-repo analyze golang/go -p 25

  # JSON output format
  go-repo analyze golang/go --format json
  go-repo analyze golang/go -f json

  # Using authentication for higher rate limits
  go-repo analyze golang/go --token ghp_xxxxxxxxxxxx --prs 50
  go-repo analyze golang/go -t ghp_xxxxxxxxxxxx -p 30 -f json`,
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&format, "format", "f", "table",
		`Output format for displaying results.
Available options:
  table - Human-readable table format with emojis (default)
  json  - Machine-readable JSON format

Examples:
  --format table  (default, shows nicely formatted table)
  --format json   (shows structured JSON data)
  -f table
  -f json`)

	analyzeCmd.Flags().IntVarP(&prs, "prs", "p", -1,
		`Number of recent pull requests to display.
Behavior:
  - Not specified: No pull requests shown (default)
  - --prs (no number): Shows 5 recent pull requests
  - --prs N: Shows N recent pull requests (max 100)

Examples:
  --prs      (shows 5 recent PRs)
  --prs 10   (shows 10 recent PRs)`)

}

func runAnalyze(cmd *cobra.Command, args []string) {
	repoURL := args[0]

	owner, repo, err := analyzer.ParseRepoURL(repoURL)
	if err != nil {
		log.Fatalf("Error parsing repository URL: %v", err)
	}

	prLimit := determinePRLimit(cmd)

	if prLimit > 100 {
		log.Fatalf("PR limit must be 100 or less")
	}

	a := analyzer.New(token)

	repoInfo, err := a.FetchRepoInfo(owner, repo)
	if err != nil {
		log.Fatalf("Error fetching repository info: %v", err)
	}

	var prInfos []*analyzer.PRInfo
	if prLimit > 0 {
		prInfos, err = a.FetchPullRequests(owner, repo, prLimit)
		if err != nil {
			log.Fatalf("Error fetching pull requests: %v", err)
		}
	}

	outputManager := output.New(format, download)

	if err := outputManager.Display(repoInfo, prInfos); err != nil {
		log.Fatalf("Error displaying output: %v", err)
	}
}

func determinePRLimit(cmd *cobra.Command) int {
	prsFlagSet := cmd.Flags().Changed("prs")

	if !prsFlagSet {
		return 0
	}

	if prs == -1 {
		return 5
	}
	return prs
}
