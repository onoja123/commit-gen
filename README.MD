# Commit-bot

A Git `commit-msg` hook powered by AI to generate concise, high-quality commit messages from your staged diff.

## Features

- Captures staged diff (`git diff --cached`)
- Sends diff to OpenAI GPT for summarization
- Automatically populates your commit message
- Written in Go for fast, standalone binary

## Prerequisites

- Go installed
- Git installed and initialized in your project
- An OpenAI API key (set as `OPENAI_API_KEY`)

## Getting Started

1. **Clone the repo**
   ```bash
   git clone https://github.com/onoja123/commit-gen.git
   cd commit-gen
   ```

2. **Build the binary**
   ```bash
   go build -o commit-msg ./cmd/commit-msg
   ```

3. **Set your OpenAI API key**
   ```bash
   export OPENAI_API_KEY="your_api_key_here"
   ```

4. **Install the Git hook**
   ```bash
   chmod +x hooks/install_hook.sh
   ./hooks/install_hook.sh
   ```

## Usage

1. Stage your changes:
   ```bash
   git add .
   ```
2. Commit as usual:
   ```bash
   git commit
   ```
3. The hook runs automatically and replaces the commit message with an AI-generated summary.

## Configuration

- **Model & parameters**: Customize in `internal/ai/openai.go`.
- **Message length**: Truncated to 72 characters by default.

## License

MIT © Onoja