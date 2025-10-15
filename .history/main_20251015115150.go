package main

import (
	"context"
	"log"
	"time"

	"ai-agent/agent"
)

func main() {
	startTime := time.Now()
	ctx := context.Background()

	// Create an AI agent with Ollama (or other LLM)
	myAgent := agent.NewAgent("llama3", agent.Calculator{})

	err := myAgent.Run(ctx, "Find the result of (4 + 5) * 2")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Execution time: %v", time.Since(startTime))
}
