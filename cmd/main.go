package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/t0mk/hotseat"
)

func main() {
	oak := os.Getenv("OPENAI_API_KEY")
	if oak == "" {
		panic("OPENAI_API_KEY environment variable not set")
	}
	os.Setenv("OPENAI_API_KEY", oak)

	funcframework.RegisterHTTPFunction("/", hotseat.GenerateTopics)

	port := "8080"
	log.Printf("Serving on port %s\n", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
