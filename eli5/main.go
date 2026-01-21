package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: eli5 <topic>")
		fmt.Fprintln(os.Stderr, "Example: eli5 \"quantum computing\"")
		os.Exit(1)
	}

	topic := os.Args[1]

	client := anthropic.NewClient()

	message, err := client.Messages.New(context.Background(), anthropic.MessageNewParams{
		Model:     anthropic.ModelClaude3_5SonnetLatest,
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			{
				Role: anthropic.MessageParamRoleUser,
				Content: anthropic.F([]anthropic.ContentBlockParamUnion{
					anthropic.TextBlockParam{
						Type: anthropic.F(anthropic.TextBlockParamTypeText),
						Text: anthropic.F(fmt.Sprintf("Explain %s like I'm 5 years old. Keep it short, fun, and use simple analogies.", topic)),
					},
				}),
			},
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, block := range message.Content {
		if block.Type == anthropic.ContentBlockTypeText {
			fmt.Println(block.Text)
		}
	}
}
