# Go Fiber Hexagonal Template

This project is a RESTful API for a simple TODO list application, built with Go and the Fiber web framework. It serves as a template and example of implementing the Hexagonal Architecture (also known as Ports and Adapters) in a Go project.

## Features

*   Get a list of all TODO items.
*   Create a new TODO item.

## Architecture

This project follows the **Hexagonal Architecture**. This architectural pattern aims to create loosely coupled application components that can be easily connected to their software environment. It isolates the core application logic from outside concerns, making the application more testable, maintainable, and adaptable to new requirements.

The main components are:

*   **Handler:** Receives HTTP requests and passes them to the service layer.
*   **Service:** Contains the core application logic and business rules.
*   **Repository:** Handles the communication with the database.

## Technologies Used

*   [Go](https://golang.org/)
*   [Fiber](https://gofiber.io/) - Web Framework
*   [PostgreSQL](https://www.postgresql.org/) - Database
*   [Docker](https://www.docker.com/) - Containerization
*   [sqlx](https://github.com/jmoiron/sqlx) - SQL extensions for Go

## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.20 or newer)
*   [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/njverse/go-fiber-hexagonal-template.git
    cd go-fiber-hexagonal-template
    ```

2.  **Create an environment file:**
    Copy the example environment file and update the variables as needed.
    ```bash
    cp .env.example .env
    ```
    Your `.env` file should look like this:
    ```
    # Server
    SERVER_PORT=8080

    # Database
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_DATABASE=go_fiber_db
    DB_SSL_MODE=disable
    ```

3.  **Start the database:**
    Run the following command to start the PostgreSQL database container.
    ```bash
    docker-compose up -d
    ```

4.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

5.  **Run the application:**
    ```bash
    go run main.go
    ```
    The server will start on the port specified in your `.env` file (default: 8080).

## API Endpoints

### Get Todos

*   **Request:**
    `GET /api/v1/todos`

*   **Description:**
    Retrieves a list of all TODO items.

*   **Response:**
    ```json
    {
        "error": false,
        "message": null,
        "data": [
            {
                "id": "some-uuid",
                "title": "My First Todo",
                "completed": false,
                "created_at": "2023-10-27T10:00:00Z"
            }
        ]
    }
    ```

### New Todo

*   **Request:**
    `POST /api/v1/todos`

*   **Description:**
    Creates a new TODO item.

*   **Body:**
    ```json
    {
        "title": "My Second Todo"
    }
    ```

*   **Response:**
    ```json
    {
        "error": false,
        "message": null,
        "data": {
            "id": "another-uuid",
            "title": "My Second Todo",
            "completed": false,
            "created_at": "2023-10-27T10:05:00Z"
        }
    }
    ```
