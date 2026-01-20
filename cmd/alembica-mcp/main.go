package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/open-and-sustainable/alembica/extraction"
	"github.com/open-and-sustainable/alembica/pricing"
	"github.com/open-and-sustainable/alembica/validation"
)

type InputJSONRequest struct {
	InputJSON     string `json:"input_json" jsonschema_description:"Full alembica input JSON" jsonschema:"required"`
	SchemaVersion string `json:"schema_version,omitempty" jsonschema_description:"Schema version override (v1 or v2)"`
}

type OutputJSONResponse struct {
	OutputJSON string     `json:"output_json,omitempty" jsonschema_description:"Result JSON string"`
	Error      *ErrorInfo `json:"error,omitempty" jsonschema_description:"Error details, if any"`
}

type ValidationResponse struct {
	Valid         bool       `json:"valid" jsonschema_description:"Whether the input validated"`
	SchemaVersion string     `json:"schema_version,omitempty" jsonschema_description:"Schema version used for validation"`
	Error         *ErrorInfo `json:"error,omitempty" jsonschema_description:"Validation error, if any"`
}

type ErrorInfo struct {
	Code    int    `json:"code" jsonschema_description:"Error code"`
	Message string `json:"message" jsonschema_description:"Error message"`
}

func main() {
	errLogger := log.New(os.Stderr, "alembica-mcp: ", log.LstdFlags)

	srv := server.NewMCPServer(
		"alembica-mcp",
		"0.2.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)

	srv.AddTool(
		mcp.NewTool(
			"alembica_validate_input",
			mcp.WithDescription("Validate alembica input JSON (schema v2 only)"),
			mcp.WithInputSchema[InputJSONRequest](),
			mcp.WithOutputSchema[ValidationResponse](),
		),
		mcp.NewStructuredToolHandler(handleValidateInput),
	)

	srv.AddTool(
		mcp.NewTool(
			"alembica_validate_output",
			mcp.WithDescription("Validate alembica output JSON (schema v2 only)"),
			mcp.WithInputSchema[InputJSONRequest](),
			mcp.WithOutputSchema[ValidationResponse](),
		),
		mcp.NewStructuredToolHandler(handleValidateOutput),
	)

	srv.AddTool(
		mcp.NewTool(
			"alembica_extract",
			mcp.WithDescription("Run alembica extraction and return output JSON (schema v2 only)"),
			mcp.WithInputSchema[InputJSONRequest](),
			mcp.WithOutputSchema[OutputJSONResponse](),
		),
		mcp.NewStructuredToolHandler(handleExtract),
	)

	srv.AddTool(
		mcp.NewTool(
			"alembica_compute_costs",
			mcp.WithDescription("Compute cost estimates for alembica input JSON (schema v2 only)"),
			mcp.WithInputSchema[InputJSONRequest](),
			mcp.WithOutputSchema[OutputJSONResponse](),
		),
		mcp.NewStructuredToolHandler(handleComputeCosts),
	)

	srv.AddTool(
		mcp.NewTool(
			"alembica_list_schemas",
			mcp.WithDescription("List supported schema versions (always v2)"),
		),
		handleListSchemas,
	)

	if err := server.ServeStdio(
		srv,
		server.WithErrorLogger(errLogger),
	); err != nil {
		errLogger.Fatalf("server error: %v", err)
	}
}

func handleValidateInput(ctx context.Context, request mcp.CallToolRequest, args InputJSONRequest) (ValidationResponse, error) {
	version, errInfo := enforceSchemaV2(args)
	if errInfo != nil {
		return ValidationResponse{
			Valid:         false,
			SchemaVersion: version,
			Error:         errInfo,
		}, nil
	}

	ctx, cancel := withTimeout(ctx)
	defer cancel()

	err := runWithTimeout(ctx, func() error {
		return validation.ValidateInput(args.InputJSON, version)
	})
	if err != nil {
		return ValidationResponse{
			Valid:         false,
			SchemaVersion: version,
			Error:         errorInfo(400, err.Error()),
		}, nil
	}

	return ValidationResponse{Valid: true, SchemaVersion: version}, nil
}

func handleValidateOutput(ctx context.Context, request mcp.CallToolRequest, args InputJSONRequest) (ValidationResponse, error) {
	version, errInfo := enforceSchemaV2(args)
	if errInfo != nil {
		return ValidationResponse{
			Valid:         false,
			SchemaVersion: version,
			Error:         errInfo,
		}, nil
	}

	ctx, cancel := withTimeout(ctx)
	defer cancel()

	err := runWithTimeout(ctx, func() error {
		return validation.ValidateOutput(args.InputJSON, version)
	})
	if err != nil {
		return ValidationResponse{
			Valid:         false,
			SchemaVersion: version,
			Error:         errorInfo(400, err.Error()),
		}, nil
	}

	return ValidationResponse{Valid: true, SchemaVersion: version}, nil
}

func handleExtract(ctx context.Context, request mcp.CallToolRequest, args InputJSONRequest) (OutputJSONResponse, error) {
	_, errInfo := enforceSchemaV2(args)
	if errInfo != nil {
		return OutputJSONResponse{Error: errInfo}, nil
	}

	ctx, cancel := withTimeout(ctx)
	defer cancel()

	var output string
	err := runWithTimeout(ctx, func() error {
		var err error
		output, err = extraction.Extract(args.InputJSON)
		return err
	})
	if err != nil {
		return OutputJSONResponse{Error: errorInfo(400, err.Error())}, nil
	}

	return OutputJSONResponse{OutputJSON: output}, nil
}

func handleComputeCosts(ctx context.Context, request mcp.CallToolRequest, args InputJSONRequest) (OutputJSONResponse, error) {
	version, errInfo := enforceSchemaV2(args)
	if errInfo != nil {
		return OutputJSONResponse{Error: errInfo}, nil
	}

	ctx, cancel := withTimeout(ctx)
	defer cancel()

	var output string
	err := runWithTimeout(ctx, func() error {
		var err error
		output, err = pricing.ComputeCosts(args.InputJSON, version)
		return err
	})
	if err != nil {
		return OutputJSONResponse{Error: errorInfo(400, err.Error())}, nil
	}

	return OutputJSONResponse{OutputJSON: output}, nil
}

func handleListSchemas(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultStructured([]string{"v2"}, "v2"), nil
}

func inferSchemaVersion(args InputJSONRequest) string {
	if args.SchemaVersion != "" {
		return args.SchemaVersion
	}

	var payload struct {
		Metadata struct {
			SchemaVersion string `json:"schemaVersion"`
		} `json:"metadata"`
	}
	if err := json.Unmarshal([]byte(args.InputJSON), &payload); err == nil {
		if payload.Metadata.SchemaVersion != "" {
			return payload.Metadata.SchemaVersion
		}
	}

	return "v1"
}

func enforceSchemaV2(args InputJSONRequest) (string, *ErrorInfo) {
	version := inferSchemaVersion(args)
	if version == "" {
		version = "v2"
	}
	if version != "v2" {
		return version, errorInfo(400, "only schema version v2 is supported by alembica-mcp")
	}
	return "v2", nil
}

func errorInfo(code int, message string) *ErrorInfo {
	return &ErrorInfo{
		Code:    code,
		Message: message,
	}
}

func withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	timeoutSeconds := os.Getenv("ALEMBICA_MCP_TIMEOUT_SECONDS")
	if timeoutSeconds == "" {
		return context.WithCancel(ctx)
	}

	seconds, err := strconv.Atoi(timeoutSeconds)
	if err != nil || seconds <= 0 {
		return context.WithCancel(ctx)
	}

	return context.WithTimeout(ctx, time.Duration(seconds)*time.Second)
}

func runWithTimeout(ctx context.Context, fn func() error) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- fn()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}
