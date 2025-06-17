# go-repo

A cobra cli tool to analyze github repositories, made this to learn Go 🐐

## Features
- Fetch repository details
- Retrieve recent pull requests
- Supports table and JSON output
- GitHub authentication
- Cross-platform

## Prerequisites

- Go 1.16 or higher
- GitHub Personal Access Token (Optional but recommended)

## Installation

```
go install github.com/w-nityammm/go-repo@latest
```
Or clone and build:
```
git clone https://github.com/w-nityammm/go-repo.git
cd go-repo
go build -o go-repo
```
## GitHub Authentication

To avoid GitHub API rate limits (60 requests/hour for unauthenticated requests), use a GitHub personal access token:

1. Create a token at: https://github.com/settings/tokens
2. Use it with go-repo in either way:
   - Pass as a flag: `--token YOUR_TOKEN` or `-t YOUR_TOKEN`
   - Set as environment variable:
     - Windows: `set GITHUB_TOKEN=YOUR_TOKEN`
     - Linux/macOS: `export GITHUB_TOKEN=YOUR_TOKEN`

## Usage

### Basic Usage

```bash
# Using owner/repo
go-repo analyze golang/go

# Using full GitHub URL
go-repo analyze https://github.com/golang/go
```

### Include Pull Requests

```bash
# Show 5 most recent pull requests
go-repo analyze golang/go --prs
go-repo analyze golang/go -p

# Show specific number of pull requests (up to 100)
go-repo analyze golang/go --prs 15
```

### Output Formats

```bash
# Table format (default)
go-repo analyze golang/go --format table
go-repo analyze golang/go -f table

# JSON format
go-repo analyze golang/go --format json
go-repo analyze golang/go -f json
```

### Help

```bash
go-repo --help
```

## Sample Output

Using `go-repo analyze golang/go -p 2`:

```
================================================================================
📦 golang/go
================================================================================
📝 The Go programming language

⭐ Stars:        128418
🍴 Forks:        18139
🐛 Open Issues:  9359
💻 Language:     Go
📅 Created:      2014-08-19
🔄 Updated:      2025-06-17

================================================================================
📋 Recent Pull Requests (2)
================================================================================
🟢 #74251: net/http: reduce allocs in CrossOriginProtection.Check
   👤 jub0bs

🔴 #74249: Victor001 hash patch 1
   👤 victor001-hash
```
## Contributing

Contributions are welcome! Few needed features:

- New flag to download output as json/pdf.
-  Show most active contributors and core maintainers.
-  Issue resolution time trends or something similar.
- Anything else you can think of!
