#!/bin/bash

# Default values
PORT=${PORT:-8086}
COUNT=${COUNT:-3}
SPECIALIZATION=${SPECIALIZATION:-""}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --port=*)
      PORT="${1#*=}"
      shift
      ;;
    --count=*)
      COUNT="${1#*=}"
      shift
      ;;
    --specialization=*)
      SPECIALIZATION="${1#*=}"
      shift
      ;;
    *)
      echo "Unknown parameter: $1"
      exit 1
      ;;
  esac
done

# Build the JSON payload
if [ -z "$SPECIALIZATION" ]; then
  JSON="{\"count\": $COUNT}"
else
  JSON="{\"count\": $COUNT, \"specialization\": \"$SPECIALIZATION\"}"
fi

# Make the curl request
echo "Sending request to http://localhost:$PORT/GenerateTopics"
echo "Payload: $JSON"
echo ""

echo $JSON

curl -X POST http://localhost:$PORT/GenerateTopics \
  -H "Content-Type: application/json" \
  -d "$JSON" | jq .

echo ""
echo "Done!"
