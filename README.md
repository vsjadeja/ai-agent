# AI Agent

An autonomous AI agent written in Go that can interact with language models and execute tools to accomplish goals. The agent features a reasoning loop that allows it to break down complex tasks, use available tools, and provide structured responses.

## Features

- 🤖 **Autonomous Reasoning**: The agent can think through problems step by step
- 🔧 **Tool Integration**: Extensible tool system for adding custom capabilities
- 🧮 **Built-in Calculator**: Includes a calculator tool for mathematical operations
- 🔄 **Iterative Problem Solving**: Uses a reasoning loop to accomplish complex goals
- 🌐 **OpenAI Compatible**: Works with OpenAI API and compatible services (like Ollama)

## Architecture

The project is structured with a clean separation of concerns:

```
ai-agent/
├── main.go              # Entry point and example usage
├── agent/
│   ├── agent.go         # Core agent implementation
│   └── tools.go         # Tool definitions (Calculator, etc.)
├── go.mod              # Go module dependencies
└── README.md           # This file
```

### Core Components

- **Agent**: The main AI agent that orchestrates thinking and tool usage
- **AIAgent Interface**: Defines the contract for agent implementations
- **Tool Interface**: Extensible interface for adding new capabilities
- **Calculator Tool**: Example tool for mathematical operations

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/vsjadeja/ai-agent.git
   cd ai-agent
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   ```bash
   # For OpenAI API
   export OPENAI_API_KEY="your-openai-api-key"
   
   # For local Ollama instance
   export OPENAI_BASE_URL="http://localhost:11434/v1"
   ```

## Usage

### Basic Example

```go
package main

import (
    "context"
    "log"
    "ai-agent/agent"
)

func main() {
    ctx := context.Background()
    
    // Create an agent with the calculator tool
    myAgent := agent.NewAgent("llama3", agent.Calculator{})
    
    // Give the agent a goal
    err := myAgent.Run(ctx, "Find the result of (4 + 5) * 2")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Running the Example

```bash
# Run with default example
go run main.go

# Or build and run
go build .
./ai-agent
```

## Configuration

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `OPENAI_API_KEY` | Your OpenAI API key | `sk-...` |
| `OPENAI_BASE_URL` | Base URL for API calls | `http://localhost:11434/v1` |

### Supported Models

The agent works with any OpenAI-compatible model:

- **OpenAI Models**: `gpt-4`, `gpt-3.5-turbo`, etc.
- **Ollama Models**: `llama3`, `mistral`, `phi3`, `codellama`, etc.
- **Other Compatible APIs**: Any service implementing OpenAI's chat completion API

## Creating Custom Tools

Tools implement the `Tool` interface and can be easily added to extend the agent's capabilities:

```go
type Tool interface {
    Name() string
    Description() string
    Execute(input string) (string, error)
}
```

### Example: Weather Tool

```go
type WeatherTool struct{}

func (w WeatherTool) Name() string {
    return "weather"
}

func (w WeatherTool) Description() string {
    return "Get current weather for a city, e.g. 'New York'"
}

func (w WeatherTool) Execute(input string) (string, error) {
    // Implement weather API call
    return fmt.Sprintf("Weather in %s: 72°F, sunny", input), nil
}

// Usage
agent := agent.NewAgent("llama3", agent.Calculator{}, WeatherTool{})
```

## API Reference

### Agent

```go
// NewAgent creates a new agent with the specified model and tools
func NewAgent(model string, tools ...Tool) *Agent

// Think sends a prompt to the language model and returns the response
func (a *Agent) Think(ctx context.Context, prompt string) (string, error)

// Run executes the agent's reasoning loop to accomplish a goal
func (a *Agent) Run(ctx context.Context, goal string) error
```

### Built-in Tools

#### Calculator

Performs basic mathematical operations:

- **Name**: `calc`
- **Supported Operations**: `+`, `-`, `*`, `/`
- **Usage**: `2+2`, `10*5`, `15/3`

## Development

### Prerequisites

- Go 1.21 or later
- OpenAI API key or local Ollama installation

### Setting up Ollama (Optional)

For local development without OpenAI API:

```bash
# Install Ollama
curl -fsSL https://ollama.com/install.sh | sh

# Pull a model
ollama pull llama3

# Start Ollama service (usually runs automatically)
ollama serve
```

### Running Tests

```bash
go test ./...
```

### Building

```bash
# Build for current platform
go build .

# Build for specific platforms
GOOS=linux GOARCH=amd64 go build .
GOOS=windows GOARCH=amd64 go build .
GOOS=darwin GOARCH=amd64 go build .
```

## Example Output

```
🧩 Goal: Find the result of (4 + 5) * 2

🤔 Agent: I need to calculate (4 + 5) * 2. Let me break this down step by step.

First, I'll calculate 4 + 5:
use:calc 4+5

Tool calc returned: 9
Continue reasoning...

🤔 Agent: Now I'll multiply 9 by 2:
use:calc 9*2

Tool calc returned: 18
Continue reasoning...

🤔 Agent: The result of (4 + 5) * 2 is 18.

✅ Final Answer: The result of (4 + 5) * 2 is 18.
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes and add tests
4. Commit your changes: `git commit -am 'Add some feature'`
5. Push to the branch: `git push origin feature-name`
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Roadmap

- [ ] Add more built-in tools (web search, file operations, etc.)
- [ ] Implement tool result caching
- [ ] Add support for tool chains and workflows
- [ ] Implement conversation memory
- [ ] Add configuration file support
- [ ] Create web interface
- [ ] Add support for structured outputs

## Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/vsjadeja/ai-agent/issues) page
2. Create a new issue with detailed information
3. Join our discussions for general questions

---

Made with ❤️ in Go
