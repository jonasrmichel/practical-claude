--- 
author: Jonas Michel
date: MMMM dd, YYYY
paging: Slide %d / %d
--- 

---

# Practical Claude
## Demos, Patterns & Pro Tips

**@jonasrmichel**

### Overview
- Setup & CLAUDE.md
- Subagents & Plan Mode
- Hooks & Permissions
- Parallel Claudes & Remote Sessions
- MCP & Verification

Press `→` for next slide, `←` for previous slide, `Ctrl+E` to execute a slide's code block, `q` to quit

---

# Working Example

Let's build a CLI tool called `eli5` that explains topics simply:

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

# Initialize with Claude (starts a new session)
claude -p "/init"
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

## Tip

Tag `@.claude` in PR comments to automatically update CLAUDE.md:

```
@.claude remember: always use context.Background()
for top-level calls in this repo
```

Claude will add this to CLAUDE.md and apply it in future sessions.

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

## Tip

Start with Sonnet, escalate to Opus for architecture.

---

# Part 2: Plan & Build

```
┌────────────────────────────────────────────────────┐
│                                                    │
│   5. Subagents                                     │
│   6. Plan Mode                                     │
│   7-8. Live Build                                  │
│   9. Hooks                                         │
│   10. Slash Commands                               │
│   11. Skills                                       │
│                                                    │
└────────────────────────────────────────────────────┘
```

---

# Subagents

Claude can spawn specialized sub-agents:

| Agent | Purpose |
|-------|---------|
| **Explore** | Search codebases, find patterns |
| **Plan** | Design implementations |
| **Bash** | Run commands |

Example:
```bash
claude -c -p "Find examples of CLI tools using the Claude API in Go"
```

Claude uses Explore agent → searches → returns findings.

## Important

Each subagent runs in its own context window with a custom system prompt, specific tool access, and independent permissions.

---

# Custom Subagents

Define specialized agents in `.claude/agents/`:

```yaml
# .claude/agents/go-expert.yaml
name: go-expert
description: Go specialist with advanced techniques
prompt: |
  You are a Go expert. Follow these principles:
  - Idiomatic Go (effective Go, Go proverbs)
  - Table-driven tests
  - Error wrapping with %w
  - Context propagation
  - Graceful shutdown patterns
  - Channel/goroutine best practices
```

## Use It

```bash
claude -c -p "Ask @go-expert to review main.go for concurrency issues"
```

---

# Proactive Subagents

Make Claude automatically delegate to your subagent:

```yaml
# .claude/agents/go-expert.yaml
name: go-expert
description: |
  Go specialist for code quality and performance.
  Use PROACTIVELY when writing, reviewing, or
  debugging Go code.
prompt: |
  You are a Go expert...
```

Key phrase: **"Use PROACTIVELY"** signals automatic delegation.

```bash
claude -c -p "Review this code for race conditions"
# Claude automatically delegates to @go-expert
```

Other trigger phrases: `"Use immediately after"`, `"MUST BE USED for"`

## Tip

If delegation isn't reliable, make the description more specific.

---

# Plan Mode

For non-trivial changes, plan first:

```bash
claude -c -p "/plan"
```

Claude will:
1. Explore the codebase
2. Design an approach
3. Write a plan file
4. Ask for your approval

## When to Use

Multi-file changes, architectural decisions, unfamiliar codebases.

<!-- DEMO: /plan to design eli5 structure -->

---

# Live Build: eli5 CLI

Let's build our CLI! Continue the session with `-c`:

```bash
claude -c -p "
  Build a Go CLI called 'eli5' that:
  1. Takes a topic as a command line argument
  2. Calls the Claude API to get an ELI5 explanation
  3. Prints the explanation to stdout

  Keep it simple - single main.go file is fine.
  Use the anthropic-sdk-go package.
"
```

<!-- DEMO: Build the CLI live -->

---

# Hooks

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
# After making changes, continue session with /commit
claude -c -p "/commit"
```

<!-- DEMO: /commit our eli5 changes -->

---

# Custom Commands

Create project-specific commands in `.claude/commands/`:

```markdown
# .claude/commands/test.md

Run all tests and show coverage report.

## Steps
1. Run `go test -cover ./...`
2. If tests fail, analyze the errors
3. Suggest fixes for any failures
```

Then invoke it:
```bash
claude -c -p "/test"
```

Commands are markdown files - Claude follows the instructions inside.

---

# Skills

Custom reusable commands:

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
```bash
claude -c -p "/eli5-test"
```

Skills = your workflow, automated.

---

# Skills: Invocation Control

Skills can be invoked **manually** (by you) or **automatically** (by Claude):

```markdown
<!-- .claude/skills/eli5-test/SKILL.md -->
~~~
name: eli5-test
description: Run eli5 with test topics
# Control who can invoke:
disable-model-invocation: false  # Claude can use (default)
user-invocable: true             # You can use (default)
~~~
```

| Setting | Effect |
|---------|--------|
| Both defaults | Claude uses when relevant, you can invoke manually |
| `disable-model-invocation: true` | Manual only (e.g., deploy scripts) |
| `user-invocable: false` | Claude only (e.g., background knowledge) |

## Tip

Write good `description` fields - Claude uses them to decide when to invoke.

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

## Tip

Start restrictive, allow as needed.

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
```bash
claude -c -p "/review-pr 123"
```

## Tip

Install the Claude Code GitHub action using `/install-github-action`.

---

# Parallel Claudes (Local)

Run multiple Claude instances on the same project:

```
Terminal 1                    Terminal 2
┌────────────────────┐       ┌────────────────────┐
│ claude -p          │       │ claude -p          │
│   "Add tests"      │       │   "Add docs"       │
└────────────────────┘       └────────────────────┘
            │                          │
            └───────────┬──────────────┘
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
claude -c --remote
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

# Teleporting: Remote → Local

Pull a remote session back to your terminal:

```bash
# List your remote sessions
claude sessions list

# Resume a remote session locally
claude sessions resume <session-id>
```

```
Remote (claude.ai/code)          Local
┌────────────────────────┐      ┌────────────────────────┐
│ Long refactor running  │      │ $ claude sessions      │
│ ...                    │  ──▶ │   resume abc123        │
│ Ready for review       │      │                        │
└────────────────────────┘      │ Resuming session...    │
                                │ Context restored ✓     │
                                └────────────────────────┘
```

Continue where Claude left off, now with full local tooling.

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

Connect Claude to external tools.

Example: fetch trending topics from r/explainlikeimfive:

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
```bash
claude -c -p "Get the top 5 posts from r/explainlikeimfive"
```
```
→ Fetching https://www.reddit.com/r/explainlikeimfive/top.json
→ 1. ELI5: Why do we dream?
→ 2. ELI5: How does WiFi actually work?
→ 3. ELI5: What causes déjà vu?
```

MCP = Claude's plugin system.

---

# MCP in Action: eli5 --top

Let's add a feature to browse r/explainlikeimfive:

```bash
claude -c -p "
  Add a --top flag to eli5 that fetches and displays
  the top K posts from r/explainlikeimfive (default K=3).
  Use the MCP fetch server to get Reddit data.
"
```

```bash
eli5 --top       # List top 3 ELI5 topics (default)
eli5 --top 5     # List top 5 ELI5 topics
```

```
$ eli5 --top
Top posts from r/explainlikeimfive:
  1. Why do we dream?
  2. How does WiFi actually work?
  3. What causes déjà vu?

Select a topic (1-3): 2

Top answer:
→ Imagine you're in a room yelling to your friend...
  WiFi is like that, but with invisible radio waves!
```

Claude + MCP fetches Reddit → user selects → shows top response.

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

## Tip

Add to hooks for automatic verification.

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
