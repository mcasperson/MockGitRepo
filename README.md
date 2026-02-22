# MockGitRepo

A boilerplate Go web application using the Gin framework.

## Prerequisites

- Go 1.21 or higher

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

## Building

Build the application:
```bash
go build -o bin/app .
```

## Running

Run the application:
```bash
./bin/app
```

Or run directly with Go:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
```bash
GET /health
```
Returns the health status of the service.

**Example:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "ok",
  "message": "Service is healthy"
}
```

### Welcome
```bash
GET /
```
Returns a welcome message.

**Example:**
```bash
curl http://localhost:8080/
```

**Response:**
```json
{
  "message": "Welcome to the Gin Web Application",
  "version": "1.0.0"
}
```

### Hello API
```bash
GET /api/hello?name={name}
```
Returns a greeting message. The `name` parameter is optional (defaults to "World").

**Example:**
```bash
curl "http://localhost:8080/api/hello?name=Copilot"
```

**Response:**
```json
{
  "greeting": "Hello, Copilot!"
}
```

## Project Structure

```
.
├── main.go          # Main application file with Gin server setup
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
└── README.md        # This file
```

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework