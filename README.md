# Go Book Store API

A simple Go-based REST API for managing books in a MySQL database using Gorilla Mux and GORM.

This project demonstrates a basic layered backend structure with:

- Router setup
- Controller layer
- Model layer
- Database configuration
- Utility helpers
- JSON request/response handling

It is designed as a beginner-friendly bookstore service where developers can create, read, update, and delete books.

---

## Table of Contents

- Overview
- Project Structure
- Tech Stack
- Prerequisites
- Setup Instructions
- Database Configuration
- Application Flow
- API Routes
- Request and Response Examples
- Models and Data Structure
- Important Notes
- Troubleshooting
- Future Improvements

---

## Overview

This project exposes a REST API for handling a book catalog. The server listens on a local port and routes incoming HTTP requests to the corresponding controller functions.

The core idea is:

1. User sends an HTTP request to the API.
2. Router matches the URL and method.
3. Controller receives the request.
4. Controller reads request JSON or parameters.
5. Model interacts with the MySQL database.
6. Response is returned as JSON.

---

## Project Structure

```text
book-store/
├── cmd/
│   └── main/
│       └── main.go
├── pkg/
│   ├── config/
│   │   └── app.go
│   ├── controllers/
│   │   └── book-controller.go
│   ├── models/
│   │   └── book.go
│   ├── routes/
│   │   └── books.routes.go
│   └── utils/
│       └── utils.go
├── go.mod
├── README.md
└── .gitignore
```

### Folder explanation

#### cmd/main/main.go

This is the application entry point. It creates the router, registers all routes, and starts the HTTP server.

#### pkg/config/app.go

This file contains the database connection logic. It initializes a global GORM DB instance and exposes helper methods to access it.

#### pkg/models/book.go

This file contains the `Book` model and methods for database operations like create, fetch all, fetch by ID, and delete.

#### pkg/controllers/book-controller.go

This layer handles HTTP requests, parses incoming input, calls model methods, and sends JSON responses.

#### pkg/routes/books.routes.go

This file defines API endpoints and attaches each route to its controller function.

#### pkg/utils/utils.go

This file contains utility methods such as request body parsing.

---

## Tech Stack

- Go
- Gorilla Mux for routing
- GORM for ORM and database management
- MySQL as the backend database
- JSON for request/response payloads

Dependencies are managed through `go.mod`.

---

## Prerequisites

Before running this project, make sure you have:

- Go installed and added to `PATH`
- MySQL server running locally or remotely
- A database created for the app
- Basic knowledge of HTTP methods: GET, POST, PUT, DELETE

Recommended:

- VS Code or any Go-supported editor
- Postman or curl for testing APIs

---

## Setup Instructions

### 1. Clone the project

```bash
git clone <your-repository-url>
cd book-store
```

### 2. Install Go dependencies

```bash
go mod tidy
```

### 3. Configure MySQL

Open the file:

```text
pkg/config/app.go
```

Update the database connection string to match your local MySQL configuration.

Example format:

```go
func Connect() {
    d, err := gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db = d
}
```

If your project is using a custom username/password and database schema, adjust the values accordingly.

### 4. Run the server

From the project root:

```bash
go run ./cmd/main
```

The service starts on:

```text
http://localhost:9010
```

---

## Database Configuration

The app uses GORM with MySQL. The database initialization happens in the `init()` function inside the model file.

### Database initialization flow

`pkg/models/book.go`:

```go
func init() {
    config.Connect()
    db = config.GetDB()
    db.AutoMigrate(&Book{})
}
```

This does two key things:

1. Connects to the configured MySQL database.
2. Creates or updates the `books` table automatically using GORM auto-migration.

### Book model

The database model is defined as:

```go
type Book struct {
    gorm.Model
    Name        string `gorm:"not null" json:"name"`
    Author      string `json:"author"`
    Publication string `json:"publications"`
}
```

#### Fields

- `ID` - auto-generated primary key from `gorm.Model`
- `CreatedAt` - auto-managed timestamp
- `UpdatedAt` - auto-managed timestamp
- `DeletedAt` - soft-delete support from `gorm.Model`
- `Name` - required book title
- `Author` - author name
- `Publication` - publication name

---

## Application Flow

### Request lifecycle

1. Request enters the HTTP server.
2. Router matches URL + method.
3. Controller method executes.
4. Controller reads body or URL variables.
5. Model interacts with `db`.
6. JSON response is returned.

### Example lifecycle: Create a book

```text
POST /book/
  -> route matches
  -> CreateBook controller executes
  -> ParseBody reads JSON from request body
  -> Book is created in database
  -> Response returns created book JSON
```

---

## API Routes

The router is defined in `pkg/routes/books.routes.go`.

### 1. Create a book

```text
POST /book/
```

#### Purpose

Creates a new book record.

#### Controller

`CreateBook`

#### Request body

```json
{
  "name": "The Alchemist",
  "author": "Paulo Coelho",
  "publications": "HarperCollins"
}
```

#### Response

Returns the created book object as JSON.

---

### 2. Get all books

```text
GET /book/
```

#### Purpose

Fetches the complete list of books.

#### Controller

`GetBook`

#### Response

Returns an array of book objects:

```json
[
  {
    "ID": 1,
    "CreatedAt": "2026-09-27T10:00:00Z",
    "UpdatedAt": "2026-09-27T10:00:00Z",
    "DeletedAt": null,
    "name": "The Alchemist",
    "author": "Paulo Coelho",
    "publications": "HarperCollins"
  }
]
```

---

### 3. Get book by ID

```text
GET /book/{bookId}
```

#### Purpose

Fetches one book by its unique ID.

#### Controller

`GetBookById`

#### Example

```text
GET /book/1
```

#### Response

```json
{
  "ID": 1,
  "CreatedAt": "2026-09-27T10:00:00Z",
  "UpdatedAt": "2026-09-27T10:00:00Z",
  "DeletedAt": null,
  "name": "The Alchemist",
  "author": "Paulo Coelho",
  "publications": "HarperCollins"
}
```

---

### 4. Update a book

```text
PUT /book/{bookId}
```

#### Purpose

Updates an existing book record.

#### Controller

`UpdateBook`

#### Example

```text
PUT /book/1
```

#### Request body

```json
{
  "name": "The Alchemist - Updated",
  "author": "Paulo Coelho",
  "publications": "HarperOne"
}
```

#### Response

Returns the updated book object.

---

### 5. Delete a book

```text
DELETE /book/{bookId}
```

#### Purpose

Deletes a book record by ID.

#### Controller

`DeleteBook`

#### Example

```text
DELETE /book/1
```

#### Response

Returns the deleted record in JSON format.

---

## Request and Response Examples

### Create a book with curl

```bash
curl -X POST http://localhost:9010/book/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Alchemist",
    "author": "Paulo Coelho",
    "publications": "HarperCollins"
  }'
```

### Get all books

```bash
curl http://localhost:9010/book/
```

### Get one book

```bash
curl http://localhost:9010/book/1
```

### Update a book

```bash
curl -X PUT http://localhost:9010/book/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Alchemist - Updated",
    "author": "Paulo Coelho",
    "publications": "HarperOne"
  }'
```

### Delete a book

```bash
curl -X DELETE http://localhost:9010/book/1
```

---

## Models and Data Structure

### Book schema

```json
{
  "ID": 1,
  "CreatedAt": "2026-09-27T10:00:00Z",
  "UpdatedAt": "2026-09-27T10:00:00Z",
  "DeletedAt": null,
  "name": "The Alchemist",
  "author": "Paulo Coelho",
  "publications": "HarperCollins"
}
```

### Meaning of each field

- `ID`: database identifier
- `CreatedAt`: when this row was created
- `UpdatedAt`: when it was last updated
- `DeletedAt`: soft-delete timestamp
- `name`: book title
- `author`: author name
- `publications`: publisher/publication house

---

## Important Notes

### 1. Router style

The project uses `gorilla/mux` and defines routes in a single registration function.

### 2. Request parsing

The utility layer reads the request body and unmarshals it into the model structure.

### 3. Database migration

The code uses `AutoMigrate`, so the table is created automatically based on the `Book` model.

### 4. JSON responses

Controllers set the response header and return the result in JSON format.

### 5. ID-based routes

The URL variable `bookId` is used to identify a specific book record.

---

## Troubleshooting

### Server not starting

Check:

- Go is installed correctly
- MySQL service is running
- Database credentials are valid
- Database name exists

### Empty or failed DB operations

Check:

- MySQL connection string in `pkg/config/app.go`
- Database user permissions
- Table creation via `AutoMigrate`

### 404 or no route found

Check:

- route path is correct
- HTTP method matches the route
- router is registered in `main.go`

### Request body not parsed

Check:

- `Content-Type` is set to `application/json`
- JSON format matches the model fields

---

## Future Improvements

This project is a nice starting point and can be extended with:

- authentication and authorization
- pagination for list endpoints
- filtering by author or publication
- validation for required fields
- error handling with proper HTTP status codes
- separate service layer for cleaner architecture
- Swagger/OpenAPI documentation
- Docker support for local development

---

## Summary

This project is a simple CRUD-style bookstore API built in Go. It follows a clean layered architecture:

- `routes` define endpoints
- `controllers` handle requests
- `models` manage database logic
- `config` manages database connection
- `utils` handle parsing helpers

It is suitable for learning Go REST APIs, MySQL integration, and GORM-based CRUD operations.

---

## Quick Start

```bash
go mod tidy
go run ./cmd/main
```

Then test endpoints using Postman or curl.

Example:

```bash
curl http://localhost:9010/book/
```

---

If you want, I can also generate:

1. a more professional `README.md` in a production style,
2. a `Postman collection` format for the API,
3. or a `Swagger/OpenAPI` version of the documentation.
