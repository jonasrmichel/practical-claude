# Practical Claude
## Demos, Patterns & Pro Tips

**@jonasrmichel**

**Topics**
- Setup & CLAUDE.md
- Plan Mode & Subagents
- Hooks & Permissions
- Parallel Claudes & Remote Sessions
- MCP & Verification

**Demo**: `eli5` - a Go CLI that explains any topic using the Claude API
```
$ eli5 "black holes"
→ A black hole is like a cosmic vacuum cleaner that's SO strong, not even light can escape!
```

Press `space` to continue, `q` to quit

---

# What We're Building

A Go CLI tool that explains topics simply:

```bash
eli5 "quantum computing"
```

```
┌──────────────────────────────────────────────────────┐
│                                                      │
│  Quantum computing is like having a magic coin       │
│  that's both heads AND tails at the same time        │
│  until you look at it...                             │
│                                                      │
└──────────────────────────────────────────────────────┘
```

We'll build this *live* while exploring Claude Code features.

---

# Part 1: Bootstrap

```
┌────────────────────────────────────────────────────┐
│                                                    │
│   1. Installation                                  │
│   2. Project Setup (/init)                         │
│   3. CLAUDE.md                                     │
│   4. Models: Opus vs Sonnet                        │
│                                                    │
└────────────────────────────────────────────────────┘
```

---

# Installation

```bash
# Install Claude Code
npm install -g @anthropic-ai/claude-code

# Verify installation
claude --version
```

```bash
# Start Claude in any directory
cd ~/projects
claude
```

Claude Code: Your AI pair programmer in the terminal.

---

# Project Setup with /init

```bash
# Create project directory
mkdir eli5 && cd eli5

# Initialize with Claude
claude
```

Then in Claude:
```
/init
```

This creates a `CLAUDE.md` file - your project's AI instruction manual.

<!-- DEMO: Run /init in the eli5 directory -->

---

# CLAUDE.md

The file that teaches Claude about *your* project:

```markdown
# CLAUDE.md

## Build Commands
go build -o eli5 .
go test ./...

## Architecture
Single-file CLI using cobra for args
Claude API for explanations

## Conventions
- Keep it simple - one file is fine
- Use table-driven tests
```

Claude reads this on every session start.

---

# Models: Opus vs Sonnet

| Aspect | Sonnet (Default) | Opus |
|--------|------------------|------|
| Speed | Fast | Slower |
| Best for | Quick tasks, iteration | Complex reasoning |
| Planning | Good | Excellent |
| Cost | Lower | Higher |

```bash
# Use Opus for complex tasks
claude --model opus

# Sonnet is default - great for most work
claude
```

**Rule of thumb**: Start with Sonnet, escalate to Opus for architecture.

---

# Part 2: Plan & Build

```
┌────────────────────────────────────────────────────┐
│                                                    │
│   5. Plan Mode                                     │
│   6. Subagents                                     │
│   7-8. Live Build                                  │
│   9. Hooks                                         │
│   10. Slash Commands                               │
│   11. Skills                                       │
│                                                    │
└────────────────────────────────────────────────────┘
```

---

# Plan Mode

For non-trivial changes, plan first:

```
/plan
```

Claude will:
1. Explore the codebase
2. Design an approach
3. Write a plan file
4. Ask for your approval

**When to use**: Multi-file changes, architectural decisions, unfamiliar codebases.

<!-- DEMO: /plan to design eli5 structure -->

---

# Subagents

Claude can spawn specialized sub-agents:

| Agent | Purpose |
|-------|---------|
| **Explore** | Search codebases, find patterns |
| **Plan** | Design implementations |
| **Bash** | Run commands |

Example prompt:
```
"Find examples of CLI tools using the Claude API in Go"
```

Claude uses Explore agent → searches → returns findings.

---

# Live Build: eli5 CLI

Let's build our CLI! The prompt:

```
Build a Go CLI called "eli5" that:
1. Takes a topic as a command line argument
2. Calls the Claude API to get an ELI5 explanation
3. Prints the explanation to stdout

Keep it simple - single main.go file is fine.
Use the anthropic-sdk-go package.
```

<!-- DEMO: Build the CLI live -->

---

# Live Build: Result

```go
///package main
///
///import (
///    "context"
///    "fmt"
///    "os"
///    "github.com/anthropics/anthropic-sdk-go"
///)
///
///func main() {
client := anthropic.NewClient()
message, _ := client.Messages.New(context.TODO(),
    anthropic.MessageNewParams{
        Model:     anthropic.ModelClaude3_5SonnetLatest,
        MaxTokens: 1024,
        Messages: []anthropic.MessageParam{{
            Role:    anthropic.MessageParamRoleUser,
            Content: "Explain " + os.Args[1] + " like I'm 5",
        }},
    })
fmt.Println(message.Content[0].Text)
///}
```

---

# Hooks: Auto-Format Code

Hooks run commands on Claude events:

```json
{
  "hooks": {
    "PostToolUse": [{
      "matcher": "Write|Edit",
      "command": ["gofmt", "-w", "$FILE_PATH"]
    }]
  }
}
```

Save to `.claude/hooks.json`

Every file write → auto-formatted with `gofmt`

<!-- DEMO: Show hook in action -->

---

# Slash Commands

Built-in productivity boosters:

| Command | Action |
|---------|--------|
| `/commit` | Smart commit with good message |
| `/review-pr` | Review a pull request |
| `/init` | Initialize CLAUDE.md |
| `/plan` | Enter plan mode |
| `/help` | Show all commands |

```bash
# After making changes
/commit
```

<!-- DEMO: /commit our eli5 changes -->

---

# Skills

Custom reusable commands (coming soon to all users):

```yaml
# .claude/skills/eli5-test.yaml
name: eli5-test
description: Run eli5 with test topics
command: |
  go build -o eli5 . && \
  ./eli5 "gravity" && \
  ./eli5 "the internet"
```

Then:
```
/eli5-test
```

Skills = your workflow, automated.

---

# Part 3: Ship & Scale

```
┌────────────────────────────────────────────────────┐
│                                                    │
│   12. Permissions                                  │
│   13. Code Review & GitHub                         │
│   14. Parallel Claudes (Local)                     │
│   15. Remote Sessions                              │
│   16. Teleporting                                  │
│                                                    │
└────────────────────────────────────────────────────┘
```

---

# Permissions

Control what Claude can do:

```json
{
  "permissions": {
    "allow": [
      "Bash(go:*)",
      "Bash(git:*)",
      "Read",
      "Write"
    ],
    "deny": [
      "Bash(rm:-rf *)",
      "Bash(curl:*)"
    ]
  }
}
```

Save to `.claude/settings.json`

**Pro tip**: Start restrictive, allow as needed.

---

# Code Review & GitHub

Claude integrates with your GitHub workflow:

**In PRs**: Tag `@claude` in comments for suggestions

**GitHub Action**: Automated review on every PR

```yaml
# .github/workflows/claude-review.yml
- uses: anthropics/claude-code-action@v1
  with:
    trigger: "pull_request"
```

**Local**:
```
/review-pr 123
```

---

# Parallel Claudes (Local)

Run multiple Claude instances on the same project:

```
Terminal 1          Terminal 2
┌──────────────┐   ┌──────────────┐
│ claude       │   │ claude       │
│              │   │              │
│ "Add tests"  │   │ "Add docs"   │
└──────────────┘   └──────────────┘
         │                 │
         └────────┬────────┘
                  ▼
            Same codebase
```

They share the filesystem but have separate contexts.

<!-- DEMO: Show parallel Claudes -->

---

# Remote Sessions

For long-running tasks, use claude.ai/code:

```bash
# Push current context to remote
claude --remote
```

**Benefits**:
- Continues when laptop closes
- Access from phone/tablet
- More compute resources

**Use case**: "Refactor this whole module" → push to remote → check later.

---

# Teleporting

Move seamlessly between local and remote:

```
Local                          Remote
┌────────────┐   teleport     ┌────────────┐
│ Started    │ ───────────▶   │ Continues  │
│ task here  │                │ running    │
└────────────┘                └────────────┘
      ▲                              │
      │        teleport back         │
      └──────────────────────────────┘
```

Context travels with you.

---

# Part 4: Power User

```
┌────────────────────────────────────────────────────┐
│                                                    │
│   17. MCP                                          │
│   18. Verification                                 │
│   19. Notifications                                │
│   20. Wrap Up                                      │
│                                                    │
└────────────────────────────────────────────────────┘
```

---

# MCP: Model Context Protocol

Connect Claude to external tools. Example: fetch trending topics from r/explainlikeimfive:

```json
{
  "mcpServers": {
    "fetch": {
      "command": "uvx",
      "args": ["mcp-server-fetch"]
    }
  }
}
```

Now Claude can fetch from the Reddit API:
```
"Get the top 5 posts from r/explainlikeimfive"
```
```
→ Fetching https://www.reddit.com/r/explainlikeimfive/top.json
→ 1. ELI5: Why do we dream?
→ 2. ELI5: How does WiFi actually work?
→ 3. ELI5: What causes déjà vu?
```

MCP = Claude's plugin system.

---

# Verification: Letting Claude Check Its Work

Add verification commands to CLAUDE.md:

```markdown
## Verification
After changes, always run:
- `go build ./...` - must compile
- `go test ./...` - tests must pass
- `go vet ./...` - no warnings
```

Claude will run these to verify its own work.

**Pro tip**: Add to hooks for automatic verification.

---

# Notifications

Never miss when Claude finishes:

**Terminal notifications with claude-notify**:
```bash
# github.com/mylee04/claude-notify
npm install -g claude-notify
```

**System bell** (built-in):
```json
{
  "notifications": {
    "onComplete": true
  }
}
```

Background a task → get notified → review.

---

# Wrap Up: What We Built

```go
///package main
///import ("context"; "fmt"; "os"; "github.com/anthropics/anthropic-sdk-go")
///func main() {
client := anthropic.NewClient()
msg, _ := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
    Model: anthropic.ModelClaude3_5SonnetLatest,
    MaxTokens: 1024,
    Messages: []anthropic.MessageParam{{
        Role: anthropic.MessageParamRoleUser,
        Content: "Explain " + os.Args[1] + " like I'm 5",
    }},
})
fmt.Println(msg.Content[0].Text)
///}
```

```bash
./eli5 "machine learning"
```

---

# Key Takeaways

1. **CLAUDE.md** - Teach Claude your project
2. **Plan Mode** - Think before coding
3. **Hooks** - Automate the repetitive
4. **Permissions** - Stay in control
5. **Parallel + Remote** - Scale your Claude
6. **Verification** - Let Claude check itself

---

# Resources

- **Docs**: docs.anthropic.com/claude-code
- **GitHub**: github.com/anthropics/claude-code
- **This talk**: github.com/jonasrmichel/practical-claude

```
┌────────────────────────────────────────────────────┐
│                                                    │
│   Questions?                                       │
│                                                    │
│   @jonasrmichel                                    │
│                                                    │
└────────────────────────────────────────────────────┘
```

Thanks for attending!
