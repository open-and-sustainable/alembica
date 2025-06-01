package logger

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

// Test default logger (Silent mode)
func TestDefaultLogger(t *testing.T) {
	// Capture original stdout and stderr
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	// Set the logger to default (which should discard logs)
	SetLogger(defaultLogger)

	// Write log messages (they should not appear)
	Info("This is an info message")
	Error("This is an error message")

	// Restore stdout and stderr
	w.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	// Read output
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	// Check that no log output was produced
	if buf.Len() != 0 {
		t.Errorf("Expected no log output, but got: %s", buf.String())
	}
}

// Test Stdout logging
func TestStdoutLogging(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := SetupLogging(Stdout, "")
	if err != nil {
		t.Fatalf("SetupLogging failed: %v", err)
	}

	Info("This is an info message")
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	if !strings.Contains(buf.String(), "This is an info message") {
		t.Errorf("Expected 'This is an info message' in log output, but got: %s", buf.String())
	}
}

// Test File logging
func TestFileLogging(t *testing.T) {
	filename := "test_log.toml"
	expectedLogFile := "test_log.log"

	err := SetupLogging(File, filename)
	if err != nil {
		t.Fatalf("SetupLogging failed: %v", err)
	}

	Info("This is a file log message")

	fileContent, err := os.ReadFile(expectedLogFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(fileContent), "This is a file log message") {
		t.Errorf("Expected 'This is a file log message' in log file, but got: %s", string(fileContent))
	}

	// Cleanup
	_ = os.Remove(expectedLogFile)
}

// Test custom logger
func TestCustomLogger(t *testing.T) {
	var buf bytes.Buffer
	customLogger := standardLogger{log.New(&buf, "", log.LstdFlags)}
	SetLogger(customLogger)

	Info("Custom logger info message")
	Error("Custom logger error message")

	output := buf.String()
	if !strings.Contains(output, "Custom logger info message") || !strings.Contains(output, "Custom logger error message") {
		t.Errorf("Expected log messages in custom logger output, but got: %s", output)
	}
}
