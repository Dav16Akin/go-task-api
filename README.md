# Go Task API

A simple and efficient RESTful API for managing tasks, built with Go. This project provides a clean and straightforward way to create, read, update, and delete tasks.

## 🚀 Features

- ✅ Create new tasks
- 📖 Retrieve all tasks or a specific task
- ✏️ Update existing tasks
- 🗑️ Delete tasks
- 💾 SQLite database for data persistence
- 🔄 Hot reload support with Air

## 🛠️ Tech Stack

- **Go** 1.25.3
- **Gin** - Fast HTTP web framework
- **GORM** - ORM library for database operations
- **SQLite** - Lightweight database
- **Air** - Live reload for Go apps

## 📋 Prerequisites

Before running this project, make sure you have:

- [Go](https://golang.org/doc/install) (version 1.25.3 or higher)
- Git

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/Dav16Akin/go-task-api.git
cd go-task-api
```

2. Install dependencies:
```bash
go mod download
```

## 🏃 Running the Application

### Standard Run

Start the server:
```bash
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

### Development Mode with Hot Reload

Install Air (if not already installed):
```bash
go install github.com/air-verse/air@latest
```

Run with Air:
```bash
air
```

The server will automatically reload when you make changes to the code.

## 📚 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v1/tasks` | Get all tasks |
| POST | `/v1/tasks` | Create a new task |
| GET | `/v1/tasks/:id` | Get a specific task |
| PUT | `/v1/tasks/:id` | Update a task |
| DELETE | `/v1/tasks/:id` | Delete a task |

## 💡 Usage Examples

### Create a Task

```bash
curl -X POST http://localhost:8080/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Complete Go tutorial",
    "completed": false
  }'
```

**Response:**
```json
{
  "result": {
    "ID": 1,
    "CreatedAt": "2024-01-08T10:00:00Z",
    "UpdatedAt": "2024-01-08T10:00:00Z",
    "DeletedAt": null,
    "title": "Learn Go",
    "description": "Complete Go tutorial",
    "completed": false
  }
}
```

### Get All Tasks

```bash
curl http://localhost:8080/v1/tasks
```

**Response:**
```json
{
  "result": [
    {
      "ID": 1,
      "CreatedAt": "2024-01-08T10:00:00Z",
      "UpdatedAt": "2024-01-08T10:00:00Z",
      "DeletedAt": null,
      "title": "Learn Go",
      "description": "Complete Go tutorial",
      "completed": false
    }
  ]
}
```

### Get a Specific Task

```bash
curl http://localhost:8080/v1/tasks/1
```

### Update a Task

```bash
curl -X PUT http://localhost:8080/v1/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Complete Go tutorial",
    "completed": true
  }'
```

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/v1/tasks/1
```

## 📁 Project Structure

```
go-task-api/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/
│   │   └── database.go       # Database configuration
│   ├── handlers/
│   │   └── task_handler.go   # HTTP request handlers
│   ├── models/
│   │   └── task.go           # Task data model
│   ├── repository/
│   │   └── task_repository.go
│   ├── requests/
│   │   └── task_requests.go  # Request validation structs
│   └── routes/
│       └── routes.go
├── .air.toml                  # Air configuration for hot reload
├── go.mod                     # Go module dependencies
├── go.sum                     # Dependency checksums
└── tasks.db                   # SQLite database file

```

## 🔧 Configuration

The application uses the following default settings:

- **Server Port:** 8080
- **Database:** SQLite (`tasks.db`)

## 📝 Task Model

A task has the following fields:

| Field | Type | Description |
|-------|------|-------------|
| ID | uint | Auto-generated unique identifier |
| CreatedAt | time.Time | Timestamp when task was created |
| UpdatedAt | time.Time | Timestamp when task was last updated |
| DeletedAt | *time.Time | Soft delete timestamp (null if not deleted) |
| title | string | Task title (required, min 3 characters) |
| description | string | Task description (max 255 characters) |
| completed | bool | Task completion status |

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📄 License

This project is open source and available for personal and educational use.

## 👤 Author

**Dav16Akin**

- GitHub: [@Dav16Akin](https://github.com/Dav16Akin)
