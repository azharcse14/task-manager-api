# Task Manager API

A small REST API for managing tasks, written in Go with SQLite. Built as a learning project.

## Features

- Create, read, update, and delete tasks (CRUD)
- Pagination with `page` and `per_page`
- Filter tasks by `done`
- Total task count in the `X-Total-Count` response header
- SQLite storage with a pure Go driver (no C compiler needed)
- No web framework: uses the standard library `net/http` router

## Tech Stack

- **Go 1.27**
- **net/http** for routing (method and path patterns like `GET /tasks/{id}`)
- **SQLite** via [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
- **encoding/json/v2** for JSON
- **log/slog** for structured logging

## Project Structure

```
task-manager-api/
├── main.go                  # Opens the database, wires everything, starts the server
├── api.http                 # Example requests for manual testing
├── ROADMAP.md               # Step-by-step build guide (in Bengali)
└── internal/
    ├── task/                # Data layer
    │   ├── task.go          # Task model
    │   ├── schema.go        # Creates the tasks table
    │   └── store.go         # SQL queries (List, GetByID, Create, Update, Delete)
    └── handler/             # HTTP layer
        ├── routes.go        # Maps URLs to handlers
        ├── task.go          # Handlers for each endpoint
        ├── query.go         # Helpers to read query parameters
        └── response.go      # Helpers to write JSON and errors
```

The code is split into two layers. `task` talks to the database and knows nothing about HTTP. `handler` deals with HTTP and never writes SQL.

## Getting Started

**Requirements:** Go 1.27 or newer. Check with `go version`.

```bash
git clone https://github.com/azharcse14/task-manager-api.git
cd task-manager-api
go run .
```

The server starts at `http://localhost:8080`. On the first run, a `tasks.db` file is created in the project folder.

## API

Base URL: `http://localhost:8080`

| Method | Path | Description | Success |
|---|---|---|---|
| GET | `/tasks` | List tasks (paginated) | 200 |
| GET | `/tasks/{id}` | Get one task | 200 |
| POST | `/tasks` | Create a task | 201 |
| PUT | `/tasks/{id}` | Update a task | 200 |
| DELETE | `/tasks/{id}` | Delete a task | 204 |

### Task Object

```json
{ "id": 1, "title": "Learn Go", "done": false }
```

| Field | Type | Notes |
|---|---|---|
| `id` | integer | Set by the server |
| `title` | string | Required, cannot be empty |
| `done` | boolean | Defaults to `false` |

JSON field names are case-sensitive. Use `"title"`, not `"Title"`.

### List Tasks

```bash
curl "http://localhost:8080/tasks?page=1&per_page=5&done=true"
```

| Query | Type | Default | Notes |
|---|---|---|---|
| `page` | integer | `1` | Values below 1 become 1 |
| `per_page` | integer | `10` | Kept between 1 and 100 |
| `done` | boolean | none | `true` or `false`. Leave it out to get all tasks |

Response:

```json
{
  "data": [
    { "id": 1, "title": "Learn Go", "done": true }
  ],
  "meta": { "page": 1, "per_page": 5, "total": 1, "total_pages": 1 }
}
```

The response also includes an `X-Total-Count` header with the total number of matching tasks.

### Get a Task

```bash
curl http://localhost:8080/tasks/1
```

Returns the task, or `404` if it does not exist.

### Create a Task

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "done": false}'
```

Returns `201` with the new task. The `Location` header points to it, for example `/tasks/1`.

### Update a Task

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go deeply", "done": true}'
```

`PUT` replaces the whole task, so send both `title` and `done`. If `done` is left out, it becomes `false`.

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

Returns `204` with no body.

### Errors

Errors are returned as JSON:

```json
{ "error": "Task not found" }
```

| Status | When |
|---|---|
| 400 | Invalid ID, invalid JSON, empty title, or a bad query value |
| 404 | Task does not exist |
| 500 | Database error |

## Manual Testing

Open `api.http` in GoLand (built-in HTTP client) or VS Code (with the REST Client extension) and run the requests one by one.

## Roadmap

- [x] CRUD endpoints
- [x] SQLite storage
- [x] Pagination and filtering
- [ ] Automated tests
- [ ] User accounts and JWT authentication
- [ ] Docker and PostgreSQL


See [LEARNING_GUIDE.md](LEARNING_GUIDE.md) for my learning progress and next steps (in Bengali).