# Deploying the IELTS Topic Generator Cloud Function

## Prerequisites

1. Google Cloud account
2. Anthropic API key
3. Google Cloud CLI installed (gcloud)
4. Go installed locally

## Setup Steps

### 1. Prepare the project

Create a new directory for your project and initialize a Go module:

```bash
mkdir hotseat
cd hotseat
go mod init github.com/t0mk/hotseat
```

### 2. Install dependencies

```bash
go get github.com/GoogleCloudPlatform/functions-framework-go/functions
go get github.com/anthropics/anthropic-sdk-go
```

### 3. Create the cloud function file

Create a file named `function.go` and paste the code from the provided example.

### 4. Deploy to Google Cloud Functions

```bash
# Set your project ID
PROJECT_ID=your-project-id

# Deploy the function
gcloud functions deploy hotseat \
  --runtime go122 \
  --trigger-http \
  --allow-unauthenticated \
  --entry-point GenerateTopics \
  --set-env-vars ANTHROPIC_API_KEY=your-anthropic-api-key \
  --project $PROJECT_ID
```

## Testing the Function

### Using curl

```bash
curl -X POST https://REGION-PROJECT_ID.cloudfunctions.net/hotseat \
  -H "Content-Type: application/json" \
  -d '{"count": 3, "specialization": "technology"}'
```
