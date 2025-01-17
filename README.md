# Library Management System

This project is a Library Management System built using Golang, MySQL, and the Fiber web framework. It provides CRUD operations for managing books and authors, including soft delete functionality.

## Features

- **Create, read, update, and delete authors**
- **Create, read, update, and delete books**
- **Soft delete functionality for authors and books**
- **Search for books by title**
- **log message if Email is changed**
- **Validation for inputs**
- **Preloading related data**

## Project Structure

.
├── main.go
├── pkg
│ ├── config
│ │ └── app.go
│ ├── controllers
│ │ ├── authorController.go
│ │ └── bookController.go
│ ├── models
│ │ ├── author.go
│ │ └── book.go
│ └── routes
│ └── routes.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md



## Setup and Installation

### Clone the repository:

```sh
git clone https://github.com/MennaOsama162/GolangAPI.git
cd library-management


Install dependencies:

Ensure you have Go and MySQL installed. Then, run:
go mod tidy

Configure the database:
Create a MySQL database named librarydb. Update the database credentials in pkg/config/app.go if necessary:

go
dsn := "root:mennaosama1682@tcp(127.0.0.1:3306)/librarydb?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true" // replace it with your connection string

Run the application:
sh
go run main.go

The server will start on http://localhost:3000.

API Endpoints

Authors
Create Author:
POST /api/authors
Request Body:
{
    "name": "Author Name",
    "email": "author@example.com"
}


Get All Authors:
GET /api/authors


Get Author by ID:
GET /api/authors/:id


Update Author:
PUT /api/authors/:id
Request Body:
{
    "name": "Updated Author Name",
    "email": "updated.author@example.com"
}


Soft Delete Author:
DELETE /api/authors/softDelete/:id


Delete Author:
DELETE /api/authors/:id


Books

Create Book:
POST /api/books
Request Body:
{
    "title": "Book Title",
    "isbn": "123-456-789",
    "published_date": "2022-02-02",
    "author_id": 1
}


Get All Books:
GET /api/books


Get Book by ID:
GET /api/books/:id


Update Book:
PUT /api/books/:id
Request Body:
{
    "title": "Updated Book Title",
    "isbn": "123-456-789",
    "published_date": "2022-02-02",
    "author_id": 1
}


Soft Delete Book:
DELETE /api/books/softDelete/:id


Delete Book:
DELETE /api/books/:id

Testing
Use Postman or any API testing tool to interact with the API endpoints.




