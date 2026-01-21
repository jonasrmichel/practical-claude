# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This repository contains a presentation and demos for "Practical Claude" - a walkthrough of Claude Code features, patterns, and pro tips.

## Repository Structure

```
practical-claude/
├── presentation.md      # Main slides (use `slides` tool to present)
├── theme.json           # Custom theme for slides
├── eli5/                # Running example: Go CLI that explains topics simply
│   ├── main.go
│   ├── go.mod
│   ├── CLAUDE.md
│   └── .claude/         # Claude Code configs for eli5
└── CLAUDE.md            # This file
```

## Commands

### Running the Presentation

```bash
# Install slides if needed
go install github.com/maaslalani/slides@latest

# Run presentation with custom theme
slides --theme theme.json presentation.md
```

### Working with eli5 Demo

```bash
cd eli5
go mod tidy           # Download dependencies
go build -o eli5 .    # Build
./eli5 "topic"        # Run
go test ./...         # Test
```

## Presentation Notes

- The presentation uses the `slides` terminal presentation tool
- Press `Ctrl+E` to execute code blocks live
- Press `space` or `→` to advance, `q` to quit
- Code blocks with `///` prefix hide boilerplate but still execute

## Demo Workflow

The presentation builds the `eli5` CLI incrementally. For live demos:
1. Start in the `eli5/` directory with a fresh state
2. Follow the slides to build features
3. Use `/commit` to commit changes at milestones
