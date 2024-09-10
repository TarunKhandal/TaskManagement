# Task Management API Server

## Overview
This project is a **Task Management API Server** written in Go, utilizing **SQLite** as the database. It provides full **CRUD operations** for managing authors, tasks, and steps, using packages like `http`, `bun` (for database connection), `gorm`, and `gookit` (for filtering and validation).

---

## Prerequisites

Ensure that you have the following prerequisites installed before setting up the project:

1. **Golang** version `>= 1.22.1`
2. **SQLite** (Optional: for manual database queries via CLI)
3. **Editor/IDE** of your choice (Recommended: **VSCode** or **GoLand**)
    - For **VSCode**, install:
      - Go Tools extension
4. **Postman** (to test API endpoints)
5. **Git**

---

## Installation

Follow these steps to install and set up the project locally:

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   ```

2. **Navigate to the project directory**:
   ```bash
   cd TaskManagement
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the server**:
   ```bash
   go run main.go
   ```

*Note: You can debug and run the program easily if you're using VSCode or other popular IDEs.*

---

## API Endpoints

The Task Management API provides endpoints for managing **Authors**, **Tasks**, and **Steps**.

### Author Endpoints

- **Create Author**  
  `POST` - `localhost:{port}/author`

- **Get All Authors**  
  `GET` - `localhost:{port}/author`

- **Get Author by ID**  
  `GET` - `localhost:{port}/author/{authorID}`

- **Update Author by ID**  
  `PUT` - `localhost:{port}/author/{authorID}`

- **Delete Author by ID**  
  `DELETE` - `localhost:{port}/author/{authorID}`

---

### Task Endpoints

- **Create Task**  
  `POST` - `localhost:{port}/tasks`

- **Get All Tasks**  
  `GET` - `localhost:{port}/tasks`

- **Get Task by ID**  
  `GET` - `localhost:{port}/tasks/{taskID}`

- **Update Task by ID**  
  `PUT` - `localhost:{port}/tasks/{taskID}`

- **Delete Task by ID**  
  `DELETE` - `localhost:{port}/tasks/{taskID}`

---

### Step Endpoints

- **Create Step**  
  `POST` - `localhost:{port}/tasks/{taskID}/steps`

- **Get All Steps for a Task**  
  `GET` - `localhost:{port}/tasks/{taskID}/steps`

- **Get Step by ID**  
  `GET` - `localhost:{port}/tasks/{taskID}/steps/{stepID}`

- **Update Step by ID**  
  `PUT` - `localhost:{port}/tasks/{taskID}/steps/{stepID}`

- **Delete Step by ID**  
  `DELETE` - `localhost:{port}/tasks/{taskID}/steps/{stepID}`

---

## Postman Documentation

You can find the Postman documentation and collection in the `postman` directory, or import the provided Postman collection file for easy testing of the API endpoints.

---

## Project Structure

```plaintext
taskmanagement
├── app
│   ├── behaviour
│   │   ├── author
│   │   ├── steps
│   │   └── task
│   ├── globals
│   │   └── dbconnections
│   ├── helpers
│   ├── http
│   │   ├── controllers
│   │   ├── middlewares
│   │   └── responses
│   └── repository
│       ├── author
│       ├── steps
│       └── task
├── configs
├── databases
├── postman
└── routes
```

---

## Development Notes

- This project uses `bun` and `gorm` for database operations. For filtering and validation, `gookit` is used.
- To manually interact with the SQLite database, you can use the SQLite CLI or your preferred SQLite database client.
