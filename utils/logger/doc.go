/*
Package logger provides a flexible logging system for the application, supporting multiple logging levels and output destinations.

Supported Logging Levels:
  - Silent: Disables all logging.
  - Stdout: Logs messages to the console (standard output).
  - File: Logs messages to a specified log file.

Core Components:
  - Logger Interface:
    - Defines `Info` and `Error` methods for logging messages.
  - standardLogger:
    - Implements `Logger` using Go’s built-in `log` package.
  - SetLogger:
    - Configures the active logger instance.
  - SetupLogging:
    - Initializes logging with the specified level and optional log file.

Features:
  - Supports **configurable logging levels**.
  - Allows logging to **stdout or a file**.
  - Uses a **default silent logger** when logging is disabled.
  - Ensures log output formatting consistency.

Example Usage:
	package main

	import (
		"log"
		"github.com/open-and-sustainable/alembica/logger"
	)

	func main() {
		// Set up logging to standard output
		err := logger.SetupLogging(logger.Stdout, "")
		if err != nil {
			log.Fatalf("Failed to set up logger: %v", err)
		}

		logger.Info("This is an informational message.")
		logger.Error("This is an error message.")
	}

*/
package logger
