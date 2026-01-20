package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"

	"github.com/open-and-sustainable/alembica/definitions"
)

type toolError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type toolOutput struct {
	OutputJSON string     `json:"output_json"`
	Error      *toolError `json:"error"`
}

func TestMCPExtractOpenAIAndDeepSeek(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping MCP live test in short mode")
	}

	openAIKey := os.Getenv("OPENAI_API_KEY")
	deepSeekKey := os.Getenv("DEEPSEEK_API_KEY")
	if openAIKey == "" || deepSeekKey == "" {
		t.Skip("Skipping MCP live test: missing OPENAI_API_KEY or DEEPSEEK_API_KEY")
	}

	repoRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("Failed to resolve repo root: %v", err)
	}
	cmdPath := filepath.Join(repoRoot, "cmd", "alembica-mcp")

	cli, err := client.NewStdioMCPClient("go", nil, "run", cmdPath)
	if err != nil {
		t.Fatalf("Failed to start MCP client: %v", err)
	}
	defer cli.Close()

	initReq := mcp.InitializeRequest{}
	initReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initReq.Params.ClientInfo = mcp.Implementation{Name: "alembica-mcp-live-test", Version: "0.1.0"}
	initReq.Params.Capabilities = mcp.ClientCapabilities{}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if _, err := cli.Initialize(ctx, initReq); err != nil {
		t.Fatalf("Failed to initialize MCP client: %v", err)
	}

	input := map[string]any{
		"metadata": map[string]any{
			"schemaVersion": "v2",
			"timestamp":     time.Now().UTC().Format(time.RFC3339),
		},
		"models": []map[string]any{
			{
				"provider":    "OpenAI",
				"api_key":     openAIKey,
				"model":       "gpt-4o-mini",
				"temperature": 0.2,
			},
			{
				"provider":    "DeepSeek",
				"api_key":     deepSeekKey,
				"model":       "deepseek-chat",
				"temperature": 0.2,
			},
		},
		"prompts": []map[string]any{
			{
				"sequenceId":     "demo",
				"sequenceNumber": 1,
				"promptContent":  "Respond with JSON: {\"status\":\"ok\"}",
			},
		},
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal input: %v", err)
	}

	callReq := mcp.CallToolRequest{}
	callReq.Params.Name = "alembica_extract"
	callReq.Params.Arguments = map[string]any{
		"input_json": string(inputBytes),
	}

	result, err := cli.CallTool(ctx, callReq)
	if err != nil {
		t.Fatalf("Tool call failed: %v", err)
	}

	outputJSON, err := extractOutputJSON(result)
	if err != nil {
		t.Fatalf("Failed to read tool output: %v", err)
	}

	var output definitions.Output
	if err := json.Unmarshal([]byte(outputJSON), &output); err != nil {
		t.Fatalf("Failed to unmarshal output JSON: %v", err)
	}

	providers := map[string]int{}
	for _, response := range output.Responses {
		providers[response.Provider]++
	}

	if providers["OpenAI"] == 0 {
		t.Fatalf("Missing OpenAI response in MCP output")
	}
	if providers["DeepSeek"] == 0 {
		t.Fatalf("Missing DeepSeek response in MCP output")
	}
}

func extractOutputJSON(result *mcp.CallToolResult) (string, error) {
	if result.StructuredContent != nil {
		var output toolOutput
		raw, err := json.Marshal(result.StructuredContent)
		if err != nil {
			return "", err
		}
		if err := json.Unmarshal(raw, &output); err != nil {
			return "", err
		}
		if output.Error != nil {
			return "", fmt.Errorf("tool error %d: %s", output.Error.Code, output.Error.Message)
		}
		if output.OutputJSON == "" {
			return "", fmt.Errorf("tool output_json is empty")
		}
		return output.OutputJSON, nil
	}

	if len(result.Content) == 0 {
		return "", fmt.Errorf("tool response has no content")
	}

	raw, err := json.Marshal(result.Content[0])
	if err != nil {
		return "", err
	}

	var text struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal(raw, &text); err != nil {
		return "", err
	}
	if text.Text == "" {
		return "", fmt.Errorf("tool response content is empty")
	}

	var output toolOutput
	if err := json.Unmarshal([]byte(text.Text), &output); err != nil {
		return "", err
	}
	if output.Error != nil {
		return "", fmt.Errorf("tool error %d: %s", output.Error.Code, output.Error.Message)
	}
	if output.OutputJSON == "" {
		return "", fmt.Errorf("tool output_json is empty")
	}
	return output.OutputJSON, nil
}
