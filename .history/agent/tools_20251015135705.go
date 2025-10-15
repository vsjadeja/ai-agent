package agent

import (
	"fmt"
	"strings"
)

type Calculator struct{}

func (c Calculator) Name() string        { return "calc" }
func (c Calculator) Description() string { return "Perform simple math operations, e.g. '2+2'" }
func (c Calculator) Execute(input string) (string, error) {
	// Remove spaces and find the operator
	input = strings.ReplaceAll(input, " ", "")

	var a, b float64
	var op string
	var opIndex int = -1

	// Find the operator position
	for i, char := range input {
		if char == '+' || char == '-' || char == '*' || char == '/' {
			// Make sure it's not a negative sign at the beginning
			if i > 0 {
				op = string(char)
				opIndex = i
				break
			}
		}
	}

	if opIndex == -1 {
		return "", fmt.Errorf("no operator found")
	}

	// Parse the numbers
	aStr := input[:opIndex]
	bStr := input[opIndex+1:]

	var err error
	if a, err = parseFloat(aStr); err != nil {
		return "", fmt.Errorf("invalid first number: %s", aStr)
	}
	if b, err = parseFloat(bStr); err != nil {
		return "", fmt.Errorf("invalid second number: %s", bStr)
	}

	switch op {
	case "+":
		return fmt.Sprintf("%.0f", a+b), nil
	case "-":
		return fmt.Sprintf("%.0f", a-b), nil
	case "*":
		return fmt.Sprintf("%.0f", a*b), nil
	case "/":
		if b == 0 {
			return "", fmt.Errorf("divide by zero")
		}
		result := a / b
		if result == float64(int64(result)) {
			return fmt.Sprintf("%.0f", result), nil
		}
		return fmt.Sprintf("%.2f", result), nil
	default:
		return "", fmt.Errorf("unknown operator: %s", op)
	}
}

func parseFloat(s string) (float64, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}
	var result float64
	n, err := fmt.Sscanf(s, "%f", &result)
	if err != nil || n != 1 {
		return 0, fmt.Errorf("failed to parse: %s", s)
	}
	return result, nil
}
