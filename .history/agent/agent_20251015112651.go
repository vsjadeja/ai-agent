package agent

import (
	"context"
	"fmt"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const OPENAI_API_KEY string = "dummy_api_key"
const OPENAI_BASE_URL string = "http://localhost:11500/v1"

type Agent struct {
	Model string
	Tools map[string]Tool
}

type AIAgent interface {
	Think(ctx context.Context, prompt string) (string, error)
	Run(ctx context.Context, goal string) error
}

type Tool interface {
	Name() string
	Description() string
	Execute(input string) (string, error)
}

// NewAgent creates an agent with registered tools
func NewAgent(model string, tools ...Tool) *Agent {
	toolMap := make(map[string]Tool)
	for _, t := range tools {
		toolMap[t.Name()] = t
	}
	return &Agent{Model: model, Tools: toolMap}
}

func (a *Agent) Think(ctx context.Context, prompt string) (string, error) {
	client := openai.NewClient(
		option.WithAPIKey(OPENAI_API_KEY),
		option.WithBaseURL(OPENAI_BASE_URL),
	)

	resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: a.Model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

// Run simulates agent reasoning loop
func (a *Agent) Run(ctx context.Context, goal string) error {
	fmt.Println("🧩 Goal:", goal)

	prompt := fmt.Sprintf(`You are an autonomous agent.
Your goal: %s

Available tools:
%s

If a tool is needed, reply with "use:<toolname> <input>"
Otherwise, reply with your final answer.`, goal, a.listTools())

	for {
		answer, err := a.Think(ctx, prompt)
		if err != nil {
			return err
		}

		answer = strings.TrimSpace(answer)
		fmt.Println("\n🤔 Agent:", answer)

		// Look for tool usage anywhere in the response
		if strings.Contains(answer, "use:") || strings.Contains(strings.ToLower(answer), "use:") {
			// Extract the tool command
			lines := strings.Split(answer, "\n")
			var toolCommand string
			for _, line := range lines {
				line = strings.TrimSpace(line)
				lineLower := strings.ToLower(line)

				// Check for "use:" at start of line
				if strings.HasPrefix(lineLower, "use:") {
					toolCommand = strings.TrimSpace(line[4:]) // Remove "use:" prefix
					break
				}
				// Check for "Use:" at start of line
				if strings.HasPrefix(line, "Use:") {
					toolCommand = strings.TrimSpace(line[4:]) // Remove "Use:" prefix
					break
				}
				// Also check for "use:" within a line
				if idx := strings.Index(lineLower, "use:"); idx != -1 {
					toolCommand = strings.TrimSpace(line[idx+4:])
					break
				}
			}

			if toolCommand != "" {
				// Parse tool and input - handle spaces in tool commands
				parts := strings.SplitN(toolCommand, " ", 2)
				if len(parts) < 2 {
					fmt.Println("⚠️  Invalid tool usage format")
					continue
				}
				toolName := strings.TrimSpace(parts[0])
				input := strings.TrimSpace(parts[1])

				fmt.Printf("🔧 Using tool: %s with input: %s\n", toolName, input)
				tool, ok := a.Tools[toolName]
				if !ok {
					fmt.Println("❌ Unknown tool:", toolName)
					continue
				}

				out, err := tool.Execute(input)
				if err != nil {
					fmt.Println("⚠️ Tool error:", err)
					continue
				}
				prompt += fmt.Sprintf("\nTool %s returned: %s\nContinue reasoning...", toolName, out)
				continue
			}
		}

		fmt.Println("✅ Final Answer:", answer)
		break
	}

	return nil
}

func (a *Agent) listTools() string {
	var out []string
	for _, t := range a.Tools {
		out = append(out, fmt.Sprintf("- %s: %s", t.Name(), t.Description()))
	}
	return strings.Join(out, "\n")
}
