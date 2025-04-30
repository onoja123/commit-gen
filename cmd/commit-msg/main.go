package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "strings"

    openai "github.com/sashabaranov/go-openai"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "usage: commit-msg <path-to-commit-msg-file>")
        os.Exit(1)
    }
    msgFile := os.Args[1]

    // 1) Get the staged diff
    out, err := exec.Command("git", "diff", "--cached", "--no-color").Output()

    if err != nil {
        fmt.Fprintf(os.Stderr, "git diff failed: %v\n", err)
        os.Exit(1)
    }

    diff := string(out)

    if strings.TrimSpace(diff) == "" {
        // no staged changes—do nothing
        os.Exit(0)
    }

    // 2) Call OpenAI
    apiKey := os.Getenv("OPENAI_API_KEY")

    if apiKey == "" {
        fmt.Fprintln(os.Stderr, "need OPENAI_API_KEY")
        os.Exit(1)
    }

    client := openai.NewClient(apiKey)
    req := openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleSystem,
                Content: "You are a helpful assistant that writes concise git commit messages.",
            },
            {
                Role:    openai.ChatMessageRoleUser,
                Content: fmt.Sprintf("Write a concise commit message (type: feat/fix/docs) for the following diff:\n\n%s", diff),
            },
        },
        MaxTokens:   64,
        Temperature: 0.2,
    }
    resp, err := client.CreateChatCompletion(context.Background(), req)

    if err != nil {
        fmt.Fprintf(os.Stderr, "OpenAI request failed: %v\n", err)
        os.Exit(1)
    }

    // 3) Extract and sanitize the message
    aiMsg := strings.TrimSpace(resp.Choices[0].Message.Content)
    lines := strings.Split(aiMsg, "\n")
    commitLine := lines[0] 
    if len(commitLine) > 72 {
        commitLine = commitLine[:72]
    }

    // 4) Overwrite the commit-msg file
    if err := ioutil.WriteFile(msgFile, []byte(commitLine+"\n"), 0644); err != nil {
        fmt.Fprintf(os.Stderr, "failed to write commit message: %v\n", err)
        os.Exit(1)
    }
}
