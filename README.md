# Go Bookstore API

A RESTful API for managing a bookstore inventory, built with Go and the Gorilla Mux router.

## Project Overview

This project implements a complete CRUD (Create, Read, Update, Delete) API for a bookstore application. It follows a structured approach with clear separation of concerns, making the codebase maintainable and scalable.

## Features

- Create, read, update, and delete books
- Search for books by ID
- Retrieve all books in the database
- RESTful API design
- Database integration with MySQL
- Environment variable configuration

## Directory Structure

```
sharukh010-go-bookstore/ 
├── go.mod 
├── go.sum 
├── cmd/ 
│   └── main/ 
│       └── main.go         # Application entry point
└── pkg/ 
    ├── config/ 
    │   └── app.go          # Database connection setup
    ├── controllers/ 
    │   └── book-controller.go  # API request handlers
    ├── models/ 
    │   └── book.go         # Book data model and DB operations
    ├── routes/ 
    │   └── bookstore-routes.go  # API route definitions
    └── utils/ 
        └── utils.go        # Common utility functions
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/book/` | Retrieve all books |
| GET | `/book/{bookId}` | Retrieve a specific book by ID |
| POST | `/book/` | Create a new book |
| PUT | `/book/{bookId}` | Update an existing book |
| DELETE | `/book/{bookId}` | Delete a book |

## Technology Stack

- Go (Golang)
- Gorilla Mux (HTTP router)
- GORM (ORM library)
- MySQL Database

## Prerequisites

- Go 1.16 or higher
- MySQL server
- Git

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/sharukh010/go-bookstore.git
   cd go-bookstore
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up environment variables:
   Create a `.env` file in the project root with the following:
   ```
   USER_DB='your_db_user'
   USER_PASSWORD='your_db_password'
   ```

4. Create a MySQL database for the application.

## Configuration

Database configuration is handled in `pkg/config/app.go`. Make sure your MySQL server is running and the credentials in your `.env` file are correct.

## Running the Application

1. Start the server:
   ```
   go run cmd/main/main.go
   ```

2. The API will be available at `http://localhost:8080`

## API Usage Examples

### Create a Book (POST)
```bash
curl -X POST http://localhost:8080/book/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Go Programming Language",
    "author": "Alan Donovan & Brian Kernighan",
    "publication": "Addison-Wesley"
  }'
```

### Get All Books (GET)
```bash
curl http://localhost:8080/book/
```

### Get Book by ID (GET)
```bash
curl http://localhost:8080/book/1
```

### Update a Book (PUT)
```bash
curl -X PUT http://localhost:8080/book/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Go Programming Language - Updated Edition",
    "author": "Alan Donovan & Brian Kernighan",
    "publication": "Addison-Wesley Professional"
  }'
```

### Delete a Book (DELETE)
```bash
curl -X DELETE http://localhost:8080/book/1
```

## Book Model

The book model contains:

- `ID`: Unique identifier
- `Name`: Book title
- `Author`: Book author
- `Publication`: Publishing company

## Development

### Adding New Features

1. Modify the Book model in `pkg/models/book.go` if needed
2. Implement new controller functions in `pkg/controllers/book-controller.go`
3. Add new routes in `pkg/routes/bookstore-routes.go`

### Extending the API

To add support for new entities (e.g., users, orders):

1. Create new model files in the `pkg/models/` directory
2. Implement corresponding controllers
3. Define new routes
4. Update database configuration if needed

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

[Sharukh010](https://github.com/sharukh010)