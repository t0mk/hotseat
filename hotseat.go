package hotseat

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GenerateTopics", GenerateTopics)
}

type TopicRequest struct {
	Count          int    `json:"count"`
	Specialization string `json:"specialization,omitempty"`
}

type Topic struct {
	Name           string   `json:"name"`
	ForbiddenWords []string `json:"forbidden_words"`
}

type TopicResponse struct {
	Topics []Topic `json:"topics"`
}

func GenerateTopics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		http.Error(w, "OPENAI_API_KEY environment variable not set", http.StatusInternalServerError)
		return
	}

	var req TopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %v", err), http.StatusBadRequest)
		return
	}

	if req.Count <= 0 {
		req.Count = 5
	}

	client := openai.NewClient(apiKey)

	randSpec := []string{
		"technology", "culture", "science", "daily life", "environment",
		"art", "business", "education", "health", "transportation",
	}[rand.Intn(10)]

	specialization := ""
	if req.Specialization != "" {
		specialization = fmt.Sprintf(" in the area of %s", req.Specialization)
	} else {
		specialization = fmt.Sprintf(" in the area of %s", randSpec)
	}

	systemPrompt := "You generate interesting, diverse IELTS speaking game topics. Return only JSON without additional text. Each topic should be unique and suitable for a 16-year-old."
	prompt := fmt.Sprintf(`Generate one IELTS "Hot Seat" speaking topic%s. Provide exactly %d forbidden words to make the topic challenging. Return JSON in format: [{"name": topic, "forbidden_words": [words]}].`, specialization, req.Count)

	resp, err := client.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model:       openai.GPT4Turbo,
		Temperature: 1,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("API request failed: %v", err), http.StatusInternalServerError)
		return
	}

	var topics []Topic
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &topics); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing API response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TopicResponse{Topics: topics})
}
