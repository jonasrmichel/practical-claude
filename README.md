<p align="center">
  <img src="assets/mascot.svg" alt="Clawd wearing a hard hat" width="150" height="150">
</p>

<h1 align="center">Practical Claude</h1>

<p align="center">
  A hands-on presentation covering demos, patterns, and pro tips for getting the most out of <a href="https://claude.ai/code">Claude Code</a>.
</p>

## What's Inside

An interactive terminal presentation that walks through the software development lifecycle while building a real CLI tool (`eli5`). Topics covered:

- **Getting Started**: Installation, `/init`, CLAUDE.md
- **Planning**: Plan mode, subagents, Opus vs Sonnet
- **Building**: Hooks, slash commands, skills
- **Shipping**: Permissions, code review, GitHub Actions
- **Scaling**: Parallel Claudes, remote sessions, teleporting
- **Power User**: MCP, verification strategies, notifications

## Requirements

- [Go](https://golang.org/) 1.21+ (for the eli5 demo)
- [slides](https://github.com/maaslalani/slides) (terminal presentation tool)
- An [Anthropic API key](https://console.anthropic.com/)

## Quick Start

```bash
# Install Claude Code (macOS/Linux/WSL)
curl -fsSL https://claude.ai/install.sh | bash

# Install slides
go install github.com/maaslalani/slides@latest

# Clone this repo
git clone https://github.com/jonasrmichel/practical-claude.git
cd practical-claude

# Run the presentation
slides presentation.md
```

## Presentation Controls

| Key | Action |
|-----|--------|
| `Space` / `→` / `j` | Next slide |
| `←` / `k` | Previous slide |
| `Ctrl+E` | Execute code block |
| `g` `g` | First slide |
| `G` | Last slide |
| `q` | Quit |

## Running the eli5 Demo

The presentation builds a Go CLI that explains topics simply:

```bash
cd eli5
go mod tidy
go build -o eli5 .
export ANTHROPIC_API_KEY=your-key-here
./eli5 "quantum computing"
```

## License

MIT
