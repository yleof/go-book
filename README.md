# Go Book API

A simple RESTful API for managing books, built with Go.

## Features

- Create a book
- List all books
- Get a book by ID
- Update a book by ID
- Delete a book by ID

## Getting Started

### Prerequisites

- Go 1.18 or newer
- Docker (for PostgreSQL)

### Run PostgreSQL with Docker

You can start a PostgreSQL container using:

```sh
docker-compose up -d
```

- user: pg
- password: pass
- db name: crud

### Install dependencies

```sh
go mod tidy
```

### Run the server

```sh
go run main.go
```

The API will start on `http://localhost:4000`.

## API Endpoints

### Add a book

```
POST /books
```

### Get all books

```
GET /books
```

### Get a book by ID

```
GET /books/{id}
```

### Update a book by ID

```
PUT /books/{id}
```

### Delete a book by ID

```
Delete /books/{id}
```

## Example Book Object

```json
{
  "id": 1,
  "title": "Book One",
  "author": "Author One",
  "desc": "This is book one"
}
```
