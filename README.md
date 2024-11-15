# Backend Intern Assignment

## Description
This backend service processes jobs related to image URL processing, where each image's perimeter is calculated after downloading. Each job can handle multiple images, and after job completion, the results are available via an API.

## Assumptions
- `store_id` should always be a string.
- Image URLs are valid and accessible.

## Setup Instructions
1. Clone this repository.
2. Navigate to the project directory.
3. Run the following commands to set up:
   - Without Docker: `go run main.go`
   - With Docker: `docker-compose up --build`
4. The API will be available at `http://localhost:8080`.

## Work Environment
- OS: macOS/Linux
- Go version: 1.18
- Editor: VS Code

## Improvements
- Add proper logging for errors.
- Implement authentication for secure API access.
