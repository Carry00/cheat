package display

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// AIRequest represents the request structure for AI API
type AIRequest struct {
	Text string `json:"text"`
}

// AIResponse represents the response structure from AI API
type AIResponse struct {
	ProcessedText string `json:"processed_text"`
}

// processWithAI sends the text to an AI service and returns the processed result
func processWithAI(text string, apiURL string, apiKey string) (string, error) {
	if apiURL == "" || apiKey == "" {
		return text, nil // Return original text if AI is not configured
	}

	// Prepare request body
	reqBody := AIRequest{Text: text}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI service returned status: %d", resp.StatusCode)
	}

	// Parse response
	var aiResp AIResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return aiResp.ProcessedText, nil
}
