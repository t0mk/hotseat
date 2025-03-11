#!/bin/bash

# Exit on any error
set -e

# Hardcoded values
FUNCTION_NAME="hotseat"
PROJECT_ID="osloveni"  # Replace with your actual project ID
REGION="europe-west12"
RUNTIME="go122"
MEMORY="256MB"
TIMEOUT="60s"
MIN_INSTANCES=0
MAX_INSTANCES=1

# Check if API key is available
if [ -z "$ANTHROPIC_API_KEY" ]; then
  echo "Error: ANTHROPIC_API_KEY environment variable is not set"
  echo "Please set it with: export ANTHROPIC_API_KEY=your_api_key"
  exit 1
fi

# Make sure Go modules are up to date
echo "Updating Go dependencies..."
go mod tidy

# Deploy the function
echo "Deploying function '$FUNCTION_NAME' to project '$PROJECT_ID' in region '$REGION'..."
gcloud functions deploy "$FUNCTION_NAME" \
  --project="$PROJECT_ID" \
  --region="$REGION" \
  --runtime="$RUNTIME" \
  --trigger-http \
  --allow-unauthenticated \
  --entry-point=GenerateTopics \
  --memory="$MEMORY" \
  --timeout="$TIMEOUT" \
  --min-instances="$MIN_INSTANCES" \
  --max-instances="$MAX_INSTANCES" \
  --set-env-vars="ANTHROPIC_API_KEY=$ANTHROPIC_API_KEY"

# Get the deployed URL
URL=$(gcloud functions describe "$FUNCTION_NAME" \
  --project="$PROJECT_ID" \
  --region="$REGION" \
  --format="value(httpsTrigger.url)")

echo ""
echo "Function deployed successfully!"
echo "URL: $URL"
echo ""
echo "Test with:"
echo "curl -X POST $URL -H \"Content-Type: application/json\" -d '{\"count\": 3}'"
