# Go Fiber Todo API

A lightweight REST API built with [Fiber](https://gofiber.io/) for managing an in-memory todo list.

## Features
- **Health Check** endpoint
- **Create Todo** (`POST /api/create-todo`)
- **Get All Todos** (`GET /api/get-todo`)
- **Update Todo** (mark as completed) (`PATCH /api/update/:id`)
- **Delete Todo** (`DELETE /api/delete/:id`)
- **Hot reload** with `air.toml` for development
- **Dockerized** for container deployment

## Tech Stack
- [Go](https://go.dev/)
- [Fiber v2](https://github.com/gofiber/fiber)
- [Air](https://github.com/cosmtrek/air) for live reloading
- Docker

---

## Installation (Local Development)
```bash
# Clone the repository
git clone <your-repo-url>
cd <repo-folder>

# Install dependencies
go mod tidy

# Run with Air (live reload)
air
