# Receipt Processor
## Overview
The Receipt Processor is a web service that processes receipts and calculates points based on predefined rules. It provides an API for submitting receipts and retrieving points awarded for each receipt.

## API Specification
The API is defined in the `api/api.yml` file. It includes the following endpoints:
- **POST /receipts/process**: Submits a receipt and returns a unique ID for the receipt.
- **GET /receipts/{id}/points**: Retrieves the points awarded for a receipt based on its ID.

## Project Structure
```
receipt-processor
├── api
│   └── api.yml
├── cmd
│   └── main.go
├── internal
│   ├── handlers
│   │   └── receipt_handler.go
│   ├── models
│   │   └── receipt.go
│   └── services
│       └── receipt_service.go
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

### Prerequisites
- Go (version 1.16 or later)
- Docker (optional, for containerization)

### Running the Application

#### Using Go
1. Clone the repository:
   ```
   git clone <repository-url>
   cd receipt-processor
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Run the application:
   ```
   go run cmd/main.go
   ```
4. The server will start on `http://localhost:8080`.

#### Using Docker
1. Build the Docker image:
   ```
   docker build -t receipt-processor .
   ```
2. Run the Docker container:
   ```
   docker run -p 8080:8080 receipt-processor
   ```
3. The server will be accessible at `http://localhost:8080`.

## Usage
- To process a receipt, send a POST request to `/receipts/process` with the receipt JSON in the body.
- To get points for a receipt, send a GET request to `/receipts/{id}/points`, replacing `{id}` with the receipt ID returned from the processing endpoint.

## License
This project is licensed for the kodurusreedharbabu with the exact propositions.