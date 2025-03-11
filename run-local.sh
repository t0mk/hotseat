#!/bin/bash

# Check if ANTHROPIC_API_KEY is set
if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo "Error: ANTHROPIC_API_KEY environment variable is not set"
    echo "Please set it with: export ANTHROPIC_API_KEY=your_api_key"
    exit 1
fi

# Default port
PORT=${PORT:-8086}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    exit 1
fi

# Create temporary main.go file for local testing
cat > main.go << 'EOF'
package main

import (
    "log"
    "os"

    "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
    _ "github.com/t0mk/hotseat/function"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8086"
    }
    
    if err := funcframework.Start(port); err != nil {
        log.Fatalf("funcframework.Start: %v\n", err)
    }
}
EOF

# Install dependencies
go mod tidy

# Run the local server
echo "Starting local server on port $PORT..."
echo "Test with: curl -X POST http://localhost:$PORT/GenerateTopics -H \"Content-Type: application/json\" -d '{\"count\": 3}'"
go run main.go
