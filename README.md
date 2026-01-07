## fe-go – Fitness & Workout Backend in Go

`fe-go` is a modern backend service written in Go that powers a simple **fitness & workout tracking** application.  
It is designed as a clean, production-style API server that handles users, workouts, and workout entries on top of a PostgreSQL database.

The project is intentionally structured to be **easy to read for recruiters, HR, and hiring managers**, while still demonstrating solid engineering practices for technical reviewers.

---

## What This Project Does

- **User management**  
  - Register and manage users with securely hashed passwords (using `bcrypt`).  
  - Basic profile data such as username, email, and bio.

- **Authentication & tokens**  
  - Token-based authentication and session handling (see `internal/tokens` and `internal/store/tokens.go`).  
  - Secure token storage using SHA-256 hashes and expiry times.

- **Workout tracking**  
  - Workouts and workout entries stored in PostgreSQL (see `internal/store/workout_store.go`).  
  - Database schema managed by SQL migrations in the `migrations/` folder (users, workouts, workout_entries, tokens).

- **HTTP API**  
  - REST-style endpoints for users and workouts (see `internal/api` and `internal/routes`).  
  - Middleware for common cross-cutting concerns (see `internal/middleware`).

In short: this repository showcases how to build a **realistic Go backend** for a fitness app, with a focus on code quality and structure rather than UI.

---

## Technologies & Tools

- **Language**: Go (Golang)
- **Web framework / routing**: `github.com/go-chi/chi/v5`
- **Database**: PostgreSQL (local instance via `docker-compose.yml`)
- **Database migrations**: `github.com/pressly/goose/v3` with SQL migration files in `migrations/`
- **Database driver / pool**: `github.com/jackc/pgx/v5/pgxpool`
- **Security & auth**:
  - `golang.org/x/crypto/bcrypt` for password hashing
  - SHA-256 for token hashing
- **Environment management**: `github.com/joho/godotenv`
- **Testing**: Go standard testing tools and `github.com/stretchr/testify` (see `workout_store_test.go`)

This mix of technologies is very typical for modern Go backend services in production.

---

## High-Level Architecture

- **`main.go`**  
  Application entry point that wires up the app, routes, and infrastructure.

- **`internal/app`**  
  Application setup, dependency wiring, and orchestration.

- **`internal/api`**  
  HTTP handlers for:
  - Users (`user_handler.go`)
  - Tokens/auth (`token_handler.go`)
  - Workouts (`workout_handler.go`)

- **`internal/routes`**  
  Route registration (mapping URLs to handlers using Chi).

- **`internal/middleware`**  
  Shared middleware for logging, authentication, etc.

- **`internal/store`**  
  Data access and domain logic for:
  - Users (`user_store.go`)
  - Workouts (`workout_store.go`)
  - Tokens (`tokens.go`)
  - Database connection and setup (`database.go`)

- **`internal/tokens`**  
  Token creation, hashing, and validation logic.

- **`migrations/`**  
  SQL files that define and evolve the database schema:
  - `00001_users.sql`, `00002_workouts.sql`, `00003_workout_entries.sql`, `0004_tokens.sql`, `00005_user_id_alter.sql`, etc.

- **`database/postgres-data`**  
  Local PostgreSQL data directory used by Docker.

This layered structure cleanly separates **API layer**, **business/domain logic**, and **persistence**.

---

## How to Run the Project Locally

**Prerequisites**

- Go installed (version compatible with `go 1.25` in `go.mod`).
- Docker and Docker Compose installed (for PostgreSQL).

**Steps**

1. **Clone the repository**

   ```bash
   git clone <this-repository-url>
   cd fe-go
   ```

2. **Start PostgreSQL via Docker Compose**

   ```bash
   docker compose up -d
   ```

3. **Install Go dependencies**

   ```bash
   go mod tidy
   ```

4. **Run database migrations** (example with `goose`)

   ```bash
   goose -dir ./migrations postgres "<your-connection-string>" up
   ```

5. **Run the application**

   ```bash
   go run ./...
   ```

6. **Access the API**

   Once running, the API will be available on the configured port (commonly `http://localhost:8080` or similar, depending on your setup).

> Replace the connection string and port above with the exact values you use in your environment variables or configuration.

---

## How to Talk About This Project (For CV / Interviews)

- **Short description**  
  “I built `fe-go`, a Go backend for a fitness & workout tracking app using PostgreSQL, Chi, and a clean layered architecture (API → services → stores).”

- **What it demonstrates**
  - Ability to design and implement a **realistic Go backend** from scratch.
  - Experience with **REST APIs**, **database design**, and **migrations**.
  - Understanding of **security best practices** (password hashing, token-based auth).
  - Familiarity with **Docker-based local development** and modern Go tooling.

This makes `fe-go` a strong portfolio project for backend / Go developer roles.

---

## Possible Next Steps / Extensions

- Add detailed API documentation (e.g. Swagger / OpenAPI).
- Implement richer validation and error responses.
- Add more workout features (plans, templates, statistics, progress charts).
- Integrate observability (structured logging, metrics, tracing via OpenTelemetry).
- Deploy to a cloud environment (e.g. containerize the app and deploy to a managed platform).

These would be natural, realistic evolutions of the current codebase for a production environment.


