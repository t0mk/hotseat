#!/bin/bash

gcloud functions deploy GenerateTopics \
  --gen2 \
  --runtime go122 \
  --region europe-west12 \
  --source . \
  --entry-point GenerateTopics \
  --trigger-http \
  --allow-unauthenticated \
  --set-env-vars OPENAI_API_KEY=$OPENAI_API_KEY

