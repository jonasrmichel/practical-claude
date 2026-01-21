# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`eli5` is a simple Go CLI that explains any topic "like I'm 5" using the Claude API.

## Build Commands

```bash
# Build
go build -o eli5 .

# Run
./eli5 "quantum computing"

# Test
go test ./...

# Format
gofmt -w .

# Vet
go vet ./...
```

## Architecture

Single-file CLI (`main.go`) that:
1. Parses command line argument (the topic)
2. Calls Claude API with an ELI5 prompt
3. Prints the response

## Dependencies

- `github.com/anthropics/anthropic-sdk-go` - Official Anthropic Go SDK

## Environment

Requires `ANTHROPIC_API_KEY` environment variable to be set.

## Conventions

- Keep it simple - single file is fine for this CLI
- Use standard library where possible
- Error messages go to stderr, output to stdout
