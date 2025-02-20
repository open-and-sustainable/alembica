package logger

import (
    "io"
    "fmt"
    "log"
    "os"
    "strings"
)

// LogLevel defines the supported levels of logging.
type LogLevel int

const (
    Silent LogLevel = iota // No logging output.
    Stdout                 // Logs messages to standard output.
    File                   // Logs messages to a file.
)

// Logger is an interface used for logging messages.
type Logger interface {
    // Info logs informational messages.
    Info(args ...interface{})
    // Error logs error messages.
    Error(args ...interface{})
}

// standardLogger is the default Logger implementation using the Go standard log package.
type standardLogger struct {
    *log.Logger
}

// Info logs informational messages.
func (l standardLogger) Info(args ...interface{}) {
    l.Println(args...)
}

// Error logs error messages.
func (l standardLogger) Error(args ...interface{}) {
    l.Println(args...)
}

// defaultLogger is a Logger that does nothing.
var defaultLogger Logger = standardLogger{log.New(io.Discard, "", log.LstdFlags)}

// activeLogger holds the currently active Logger.
var activeLogger Logger = defaultLogger

// SetLogger sets the active Logger for use in the application.
//
// Parameters:
//   - l: The Logger instance to use.
func SetLogger(l Logger) {
    activeLogger = l
}

// Info logs an informational message using the active logger.
func Info(args ...interface{}) {
    activeLogger.Info(args...)
}

// Error logs an error message using the active logger.
func Error(args ...interface{}) {
    activeLogger.Error(args...)
}

// SetupLogging configures the logging output based on the specified log level.
// If File logging is chosen, it will write logs to a file.
//
// Parameters:
//   - level: The desired log level (Silent, Stdout, File).
//   - filename: The filename to use when logging to a file (if applicable).
//
// Returns:
//   - An error if file logging fails to initialize.
func SetupLogging(level LogLevel, filename string) error {
    var logOutput io.Writer
    switch level {
    case Silent:
        logOutput = io.Discard
    case Stdout:
        logOutput = os.Stdout
    case File:
        logname := strings.TrimSuffix(filename, ".toml") + ".log"
        logFile, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            return fmt.Errorf("failed to open log file: %v", err)
        }
        logOutput = logFile
    default:
        logOutput = os.Stdout // Default to stdout if no valid option is provided
    }
    SetLogger(standardLogger{log.New(logOutput, "", log.LstdFlags)})
    return nil
}
