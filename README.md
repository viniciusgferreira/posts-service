# Posts Microservice (service-posts)

This repository houses the official source code for the Posts Microservice, a core component of the blog platform architecture. The primary responsibility of this service is to manage the complete lifecycle of post entities, encompassing all Create, Read, Update, and Delete (CRUD) operations.

## Core Technologies

* **Language:** Go (Golang)
* **Web Framework:** Gin
* **Database:** MongoDB
* **Containerization:** Docker

## Local Deployment and Execution

To deploy and run the service within a local environment, the following steps should be executed.

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

## API Specification

The service exposes the following RESTful endpoints for interaction.

| Method   | Path                  | Description                                                              |
| :------- | :-------------------- | :----------------------------------------------------------------------- |
| `POST`   | `/api/posts`       | Facilitates the creation of a new post entity.                           |
| `GET`    | `/api/posts/:id`   | Retrieves a single post entity, identified by its unique ID.             |
| `GET`    | `/api/posts`       | Returns a paginated list of all post entities.                           |
| `PUT`    | `/api/posts/:id`   | Updates the content and attributes of a pre-existing post entity.        |
| `DELETE` | `/api/posts/:id`   | Permanently removes a post entity from the system.                       |
| `GET`    | `/health-check`       | Provides an endpoint to verify the operational status of the service.    |
