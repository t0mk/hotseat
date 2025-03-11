// Package ieltstopics provides a Google Cloud Function that generates IELTS Hot Seat topics
package ieltstopics

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	anthropic "github.com/anthropics/anthropic-sdk-go"
)

func init() {
	functions.HTTP("GenerateTopics", generateTopics)
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

func generateTopics(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for browser requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	// Get API key from environment variable
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		http.Error(w, "ANTHROPIC_API_KEY environment variable not set", http.StatusInternalServerError)
		return
	}

	var req TopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %v", err), http.StatusBadRequest)
		return
	}

	// Set default count if not specified
	if req.Count <= 0 {
		req.Count = 5
	}

	// Create Anthropic client
	client := anthropic.NewClient()

	randSpec := []string{
		"technology", "culture", "science", "daily life", "environment",
		"art", "business", "education", "health", "transportation",
		"architecture", "astronomy", "agriculture", "archaeology", "aviation",
		"biology", "botany", "chemistry", "cooking", "crafts",
		"cryptocurrency", "cybersecurity", "dance", "design", "economics",
		"energy", "engineering", "ethics", "fashion", "festivals",
		"film", "finance", "folklore", "food", "forestry",
		"furniture", "games", "gardening", "genealogy", "geography",
		"geology", "history", "hobbies", "holidays", "home decoration",
		"humor", "hydrology", "industrial processes", "insects", "interior design",
		"journalism", "languages", "law", "linguistics", "literature",
		"logistics", "maritime", "marketing", "materials science", "mathematics",
		"mechanics", "media", "medicine", "metallurgy", "meteorology",
		"microbiology", "military", "minerals", "music", "mythology",
		"nanotechnology", "neuroscience", "nutrition", "oceanography", "paleontology",
		"pets", "pharmacy", "philosophy", "photography", "physics",
		"politics", "psychology", "public speaking", "publishing", "recycling",
		"religion", "robotics", "safety", "social media", "sociology",
		"sports", "statistics", "sustainable living", "textiles", "theater",
		"tools", "tourism", "toys", "traditional practices", "urban planning",
		"veterinary science", "waste management", "water conservation", "wildlife", "zoology",
	}[rand.Intn(100)]

	// Create the prompt
	specialization := ""
	if req.Specialization != "" {
		specialization = fmt.Sprintf(" in the area of %s", req.Specialization)
	} else {
		specialization = fmt.Sprintf(" in the area of %s", randSpec)
	}

	systemPrompt := "You are a helpful API that generates interesting, diverse IELTS speaking game topics. Return only valid JSON with no additional text. Each topic should be unique and different from common examples like Urban Beekeeping. A 16 years old should be able to describe the topic."
	prompt := fmt.Sprintf(`Generate a topic for an IELTS speaking game called "Hot Seat"%s.
Also provide exactly %d forbidden words that would make the topic hard to describe if they cannot be used.
Format your response as a JSON array of objects with "name" and "forbidden_words" properties. The array should have only 1 object. Be carefult not to put any backtics in the response.
Example (if count was 5):
[
  {
    "name": "Urban Beekeeping",
    "forbidden_words": ["honey", "hive", "sting", "pollinate", "queen"]
  }
]`, specialization, req.Count)
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Temperature: anthropic.F(1.),
		Model:       anthropic.F(anthropic.ModelClaude3_7SonnetLatest),
		MaxTokens:   anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		}),
		System: anthropic.F([]anthropic.TextBlockParam{
			anthropic.NewTextBlock(systemPrompt),
		}),
	})
	if err != nil {
		panic(err.Error())
	}

	// Extract JSON content from response
	content := message.Content[0].Text
	// Parse JSON response
	var topics []Topic
	if err := json.Unmarshal([]byte(content), &topics); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing API response from %s: %v", content, err), http.StatusInternalServerError)
		return
	}

	// Return the topics
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TopicResponse{Topics: topics})
}
