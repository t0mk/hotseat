# Hot Seat - IELTS Speaking Topic Generator

A web application that generates challenging IELTS "Hot Seat" speaking topics with forbidden words, powered by OpenAI's GPT-4 model and deployed as a Google Cloud Function.

## Features

- Generate IELTS speaking topics with customizable difficulty
- Specify the number of forbidden words to make topics more challenging
- Filter topics by specialization/category (technology, environment, education, etc.)
- Responsive web interface with printable topic cards
- Serverless backend using Google Cloud Functions

## Technology Stack

- **Backend**:
  - Go 1.22
  - Google Cloud Functions
  - OpenAI GPT-4 API
  - Functions Framework for Go

- **Frontend**:
  - HTML5
  - CSS3
  - Vanilla JavaScript

## Prerequisites

- Go 1.22 or later
- Google Cloud SDK
- OpenAI API key
- Git

## Installation & Deployment

### Local Development

1. Clone the repository:
   ```
   git clone https://github.com/t0mk/hotseat.git
   cd hotseat
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up your OpenAI API key:
   ```
   export OPENAI_API_KEY=your_api_key_here
   ```

### Deployment to Google Cloud Functions

1. Make sure you have Google Cloud SDK installed and are logged in:
   ```
   gcloud auth login
   ```

2. Deploy using the provided script:
   ```
   ./deploy.sh
   ```

   This script will deploy the function to Google Cloud Functions in the europe-west12 region.

3. After deployment, you'll receive a URL for your function that can be used in the frontend.

### Setting Up the Frontend

1. Update the `API_URL` in `index.html` with your deployed function URL:
   ```javascript
   const API_URL = 'your_function_url_here';
   ```

2. Host the `index.html` file on any web server or static hosting service.

## Usage

1. Open the web application in your browser.
2. Customize the generation parameters:
   - Set the number of forbidden words (3-15)
   - Optionally specify a specialization area (e.g., technology, environment)
3. Click "Generate a Topic" to create a new speaking topic.
4. Use the "Print Cards" button to print the generated topic cards.

The application provides speaking topics with a list of words that must not be used while discussing the topic, creating a challenging speaking exercise.

## Development

### Project Structure

- `hotseat.go` - Main application code for the Google Cloud Function
- `go.mod` & `go.sum` - Go module dependencies
- `deploy.sh` - Deployment script for Google Cloud Functions
- `index.html` - Frontend interface
- `/cmd` - Code to run the function locally

### Local Testing

To test the function locally, you can use the Functions Framework for Go:

```
go run cmd/local/main.go
```

### API Format

#### Request:
```json
{
  "count": 5,
  "specialization": "technology"
}
```

#### Response:
```json
{
  "topics": [
    {
      "name": "Smart Home Devices",
      "forbidden_words": ["internet", "control", "assistant", "connect", "voice"]
    }
  ]
}
```

### Environment Variables

- `OPENAI_API_KEY` - Required for making requests to the OpenAI API

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]

