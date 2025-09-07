package display

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/cheat/cheat/internal/config"
)

// ChatMessage represents a message in the conversation
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AIRequest represents the request structure for OpenAI API
type AIRequest struct {
	Model     string        `json:"model"`
	Messages  []ChatMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens,omitempty"`
}

// Choice represents a response choice from OpenAI
type Choice struct {
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

// AIResponse represents the response structure from OpenAI API
type AIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
}

// processWithAI sends the text to an AI service and returns the processed result
func processWithAI(text string, conf config.Config) (string, error) {
	if !conf.AIEnabled || conf.AIURL == "" || conf.AIKey == "" {
		return text, nil // Return original text if AI is not configured
	}

	model := conf.AIModel
	if model == "" {
		model = "gpt-3.5-turbo" // Default model
	}

	systemPrompt := conf.AISystemPrompt
	if systemPrompt == "" {
		systemPrompt = "你是一个AI助手，负责处理cheat命令的输出。请：\n1. 增强内容的可读性\n2. 为专业术语添加简短解释\n3. 保持原始命令和示例的准确性"
	}

	// Prepare request body
	reqBody := AIRequest{
		Model: model,
		Messages: []ChatMessage{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: text,
			},
		},
	}
	if conf.AIMaxTokens > 0 {
		reqBody.MaxTokens = conf.AIMaxTokens
	}
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", conf.AIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.AIKey))

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

	// Check if we have any choices in the response
	if len(aiResp.Choices) == 0 {
		return "", fmt.Errorf("no response from AI service")
	}

	// Return the content from the first choice
	return aiResp.Choices[0].Message.Content, nil
}
