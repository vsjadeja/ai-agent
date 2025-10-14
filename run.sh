#!/bin/bash

# Set environment variables for Ollama
export OPENAI_BASE_URL="http://localhost:11434/v1"
export OPENAI_API_KEY="dummy-key"

# Run the AI agent
echo "🚀 Starting AI Agent with Ollama..."
echo "📍 Using Ollama at: $OPENAI_BASE_URL"
echo ""

go run main.go