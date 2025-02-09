package alembica

import (
    "encoding/json"
    "fmt"
	"github.com/open-and-sustainable/alembica/utils/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/open-and-sustainable/alembica/llm/tokens"
	"sync"
	"time"
)


// Global variable to store the timestamps of requests
var requestTimestamps []time.Time
var mutex sync.Mutex

func Extract(inputJSON string) (string, error) {
    // Unmarshal JSON input to Input struct
    var inputData definitions.Input
    err := json.Unmarshal([]byte(inputJSON), &inputData)
    if err != nil {
		msg := fmt.Sprintf("error parsing input JSON: %v", err)
		logger.Error(msg)
		return "", err
    }
    // Simulate processing logic and generate output
    outputData := definitions.Output{
        Metadata: definitions.OutputMetadata{
            SchemaVersion: inputData.Metadata.SchemaVersion,
        },
        Responses: make([]definitions.Response, len(inputData.Prompts)),
    }

    // Example processing logic: just echoing back input in a simplified form
    for i, prompt := range inputData.Prompts {
        outputData.Responses[i] = definitions.Response{
            SequenceID:    prompt.SequenceID,
            SequenceNumber: prompt.SequenceNumber,
            ModelResponse: "Response to " + prompt.PromptContent,
        }
    }

    // Marshal output data to JSON
    outputJSON, err := json.Marshal(outputData)
    if err != nil {
        return "", fmt.Errorf("error generating output JSON: %v", err)
    }

    return string(outputJSON), nil
}


func AssessCost () error {
	return nil
}

func Validate () error {
	return nil
}

// Method that returns the number of seconds to wait to respect TPM limits
func getWaitTime(prompt string, llm definitions.Model) int {
	// Locking to ensure thread-safety when accessing the requestTimestamps slice
	mutex.Lock()
	defer mutex.Unlock()

	// Clean up old timestamps (older than 60 seconds)
	now := time.Now()
	cutoff := now.Add(-60 * time.Second)
	validTimestamps := []time.Time{}
	for _, timestamp := range requestTimestamps {
		if timestamp.After(cutoff) {
			validTimestamps = append(validTimestamps, timestamp)
		}
	}
	requestTimestamps = validTimestamps

	// Add the current request timestamp
	requestTimestamps = append(requestTimestamps, now)

	// Get the current number of requests in the last 60 seconds
	numRequests := len(requestTimestamps)

	// Calculate the time to wait until the next minute
	remainingSeconds := 60 - now.Second()
	// Analyze TPM limits
	tpm_wait_seconds := 0
	// Get the TPM limit from the configuration
	tpmLimit := llm.TPMLimit
	if tpmLimit > 0 {
		// Get the number of tokens from the prompt
		counter := tokens.RealTokenCounter{}
		tokens := counter.GetNumTokensFromPrompt(prompt, llm.Provider, llm.Model, llm.APIKey)
		tpm_wait_seconds = remainingSeconds
		// Calculate the number of tokens per second allowed
		tokensPerSecond := float64(tpmLimit) / 60.0
		// Calculate the required wait time in seconds to not exceed TPM limit
		requiredWaitTime := float64(tokens) / tokensPerSecond
		// Calculate the seconds to the next minute
		secondsToMinute := 0
		if int(requiredWaitTime) > 60 {
			secondsToMinute = 60 - int(requiredWaitTime)%60
		}
		// If required wait time is more than remaining seconds in the current minute, wait until next minute
		if requiredWaitTime > float64(remainingSeconds) {
			tpm_wait_seconds = remainingSeconds + int(requiredWaitTime) + secondsToMinute
		}
		// Otherwise, calculate the wait time based on tokens used
	}
	// Analyze RPM limits
	rpm_wait_seconds := 0
	rpmLimit := llm.RPMLimit
	if rpmLimit > 0 {
		// If the number of requests risks to exceed the RPM limit, we need to wait
		if numRequests >= int(rpmLimit-1) {
			rpm_wait_seconds = remainingSeconds
		}
	}

	// Return the maximum of tpm_wait_seconds and rpm_wait_seconds
	if tpm_wait_seconds > rpm_wait_seconds {
		return tpm_wait_seconds
	} else {
		return rpm_wait_seconds
	}
}

func waitWithStatus(waitTime int) {
	ticker := time.NewTicker(1 * time.Second) // Ticks every second
	defer ticker.Stop()
	remainingTime := waitTime
	for range ticker.C {
		// Print the status only when the remaining time modulo 5 equals 0
		if remainingTime%5 == 0 {
			logger.Info("Waiting... %d seconds remaining\n", remainingTime)
		}
		remainingTime--
		// Break the loop when no time is left
		if remainingTime <= 0 {
			logger.Info("Wait completed.")
			break
		}
	}
}


