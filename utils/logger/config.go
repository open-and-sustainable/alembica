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
    Silent LogLevel = iota
    Stdout
    File
)

// Logger interface used throughout the library to abstract the actual logging.
type Logger interface {
    Info(args ...interface{})
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
var defaultLogger = standardLogger{log.New(io.Discard, "", 0)}

// activeLogger holds the currently active Logger.
var activeLogger Logger = defaultLogger

// SetLogger sets the active Logger.
func SetLogger(l Logger) {
    activeLogger = l
}

// Info delegates an info message to the active logger.
func Info(args ...interface{}) {
    activeLogger.Info(args...)
}

// Error delegates an error message to the active logger.
func Error(args ...interface{}) {
    activeLogger.Error(args...)
}

// SetupLogging configures the logging output based on the specified log level and optional file name for file logging.
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
