package model

import (
	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/open-and-sustainable/alembica/llm/tokens"
	"sync"
	"time"
)

// Global variables for tracking request timestamps to enforce rate limits.
var requestTimestamps []time.Time
var mutex sync.Mutex

// Wait enforces rate limits by delaying execution based on token per minute (TPM) and request per minute (RPM) constraints.
//
// Parameters:
//   - prompt: The text prompt being processed.
//   - llm: The model configuration containing rate limits.
func Wait(prompt string, llm definitions.Model) {
	waitTime := getWaitTime(prompt, llm)
	waitWithStatus(waitTime)
}

// getWaitTime calculates the required wait time in seconds based on TPM and RPM limits.
//
// Parameters:
//   - prompt: The text prompt being processed.
//   - llm: The model configuration containing rate limits.
//
// Returns:
//   - The number of seconds to wait before the next request.
func getWaitTime(prompt string, llm definitions.Model) int {
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
	remainingSeconds := 60 - now.Second()
	
	// Analyze TPM limits
	tpm_wait_seconds := 0
	tpmLimit := llm.TPMLimit
	if tpmLimit > 0 {
		counter := tokens.RealTokenCounter{}
		tokenCount := counter.GetNumTokensFromPrompt(prompt, llm.Provider, llm.Model, llm.APIKey)
		tokensPerSecond := float64(tpmLimit) / 60.0
		requiredWaitTime := float64(tokenCount) / tokensPerSecond
		if requiredWaitTime > float64(remainingSeconds) {
			tpm_wait_seconds = remainingSeconds + int(requiredWaitTime)
		}
	}
	
	// Analyze RPM limits
	rpm_wait_seconds := 0
	rpmLimit := llm.RPMLimit
	if rpmLimit > 0 && numRequests >= int(rpmLimit-1) {
		rpm_wait_seconds = remainingSeconds
	}

	// Return the maximum required wait time
	if tpm_wait_seconds > rpm_wait_seconds {
		return tpm_wait_seconds
	}
	return rpm_wait_seconds
}

// waitWithStatus enforces a waiting period, displaying status updates every 5 seconds.
//
// Parameters:
//   - waitTime: The number of seconds to wait.
func waitWithStatus(waitTime int) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	remainingTime := waitTime
	for range ticker.C {
		if remainingTime%5 == 0 {
			logger.Info("Waiting... %d seconds remaining\n", remainingTime)
		}
		remainingTime--
		if remainingTime <= 0 {
			logger.Info("Wait completed.")
			break
		}
	}
}
