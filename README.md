# Posts Microservice (service-posts)

This repository houses the official source code for the Posts Microservice, a core component of the blog platform architecture. The primary responsibility of this service is to manage the complete lifecycle of post entities, encompassing all Create, Read, Update, and Delete (CRUD) operations.

## Core Technologies

* **Language:** Go (Golang)
* **Web Framework:** Gin
* **Database:** MongoDB
* **Containerization:** Docker

## Features

- **Graceful Shutdown**: Properly handles SIGINT and SIGTERM signals for clean shutdown
- **Health Check**: Built-in health check endpoint
- **CORS Support**: Cross-origin resource sharing enabled
- **Structured Logging**: JSON-formatted logging with logrus
- **Clean Architecture**: Follows clean architecture principles

## Local Deployment and Execution

### Prerequisites
- Go 1.21 or higher

### Quick Start

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the application:**
   ```bash
   go run cmd/main.go
   ```

   The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable:
   ```bash
   PORT=3000 go run cmd/main.go
   ```

3. **Test the health check:**
   ```bash
   curl http://localhost:8080/health
   ```

### Environment Variables

- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin mode (debug/release, default: release)

### Docker Deployment

1.  **Container Image Construction:**
    First, construct the Docker container image using the provided Dockerfile.

    ```sh
    docker build -t service-posts .
    ```

2.  **Container Instantiation:**
    Subsequently, instantiate the container from the newly created image. It is imperative that the required environment variables, specifically the database connection string, are provided at runtime.

    ```sh
    docker run -p 8080:8080 \
      -e MONGODB_URI="mongodb://user:password@host:port/dbname" \
      service-posts
    ```

### Graceful Shutdown

The application supports graceful shutdown. When you send a SIGINT (Ctrl+C) or SIGTERM signal, the server will:

1. Stop accepting new requests
2. Wait for existing requests to complete (up to 30 seconds)
3. Shutdown cleanly

## API Specification

The service exposes the following RESTful endpoints for interaction.

### Health Check
| Method   | Path                  | Description                                                              |
| :------- | :-------------------- | :----------------------------------------------------------------------- |
| `GET`    | `/health`             | Provides an endpoint to verify the operational status of the service.    |

### Posts API (v1)
| Method   | Path                  | Description                                                              |
| :------- | :-------------------- | :----------------------------------------------------------------------- |
| `POST`   | `/api/v1/posts`       | Facilitates the creation of a new post entity.                           |
| `GET`    | `/api/v1/posts/:id`   | Retrieves a single post entity, identified by its unique ID.             |
| `GET`    | `/api/v1/posts`       | Returns a paginated list of all post entities.                           |
| `PUT`    | `/api/v1/posts/:id`   | Updates the content and attributes of a pre-existing post entity.        |
| `DELETE` | `/api/v1/posts/:id`   | Permanently removes a post entity from the system.                       |

### Example Health Check Response
```json
{
  "status": "healthy",
  "service": "posts-service"
}
```
