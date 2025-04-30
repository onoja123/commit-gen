#!/bin/sh

# Determine the hooks directory in the current Git repo
HOOKS_DIR=$(git rev-parse --git-dir)/hooks

# Path to the built Go binary (assumes you're in the repo root)
BIN_PATH="$(pwd)/smart-commit-msg"

# Write the commit-msg hook
cat > "$HOOKS_DIR/commit-msg" << 'EOF'
#!/bin/sh
# Forward the hook argument (the path to .git/COMMIT_EDITMSG) to our binary
exec "$BIN_PATH" "$1"
EOF

# Make the hook executable
chmod +x "$HOOKS_DIR/commit-msg"

echo "✅ Installed smart-commit-bot commit-msg hook."
