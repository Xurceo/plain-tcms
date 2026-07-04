[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

# Plain TCMS

A lightweight Test Case Management System with a Go/Gin backend and React frontend.

## Tech Stack

- **Backend**: Go, Gin, GORM, PostgreSQL
- **Frontend**: React 19, TypeScript, Vite, Tailwind CSS
- **Docs**: Swagger / OpenAPI 2.0

## Getting Started

### Prerequisites

- Go 1.26+
- Node.js 22+
- PostgreSQL 16+

### Database

Create the database and run the schema migration:

```bash
createdb tcms
psql -d tcms -f backend/db/migrations/001_init.sql
```

Optionally load seed data:

```bash
psql -d tcms -f backend/db/migrations/test_rows.sql
```

### Backend

```bash
cd backend

# Copy and configure environment
cp .env.example .env

# Run the server (auto-installs deps via go.mod)
go run .
```

Available at `http://localhost:8080`. Swagger UI at `http://localhost:8080/swagger/index.html`.

### Generate Swagger docs

Swagger docs only generate after adding comments to endpoints and main
app like so\:

```go
// @title Plain-TCMS API
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

```go
// GetProjects godoc
// @Summary Get all projects
// @Tags Projects
// @Produce json
// @Success 200 {array} entities.Project
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects [get]
func GetProjects() {}
```

And then generate swagger docs out of comments:

```bash
swag init
```

### Frontend

```bash
cd frontend
pnpm install
pnpm dev
```

Available at `http://localhost:5173`.

## Environment Variables

### Backend

| Variable       | Default                                        | Description           |
|----------------|------------------------------------------------|-----------------------|
| `PORT`         | `8080`                                         | HTTP server port      |
| `DATABASE_URL` | `postgresql://tcms:tcms@localhost:5432/tcms`   | PostgreSQL connection |

### Frontend

| Variable        | Default                          | Description          |
|-----------------|----------------------------------|----------------------|
| `VITE_API_URL`  | `http://localhost:8080/api/v1`   | Backend API base URL |

## API

All routes are under `/api/v1`.

### Organizations

| Method | Path                         | Description            |
|--------|------------------------------|------------------------|
| GET    | `/organizations`             | List all organizations |
| GET    | `/organizations/:id`         | Get organization by ID |
| POST   | `/organizations`             | Create organization    |
| DELETE | `/organizations/:id`         | Delete organization    |

### Projects

| Method | Path                          | Description          |
|--------|-------------------------------|----------------------|
| GET    | `/projects`                   | List all projects    |
| GET    | `/projects/:id`               | Get project by ID    |
| DELETE | `/projects/:id`               | Delete project       |
| GET    | `/organizations/:id/projects` | List projects by org |
| POST   | `/organizations/:id/projects` | Create project       |

### Test Cases

| Method | Path                            | Description                |
|--------|---------------------------------|----------------------------|
| GET    | `/projects/:id/test-cases`      | List test cases by project |
| POST   | `/projects/:id/test-cases`      | Create test case           |

## Project Structure

```
backend/
├── db/
│   ├── db.go              # GORM connection
│   └── migrations/        # SQL schema + seed data
├── entities/              # GORM model structs
├── repository/            # Data access layer
├── endpoints/             # HTTP handlers
├── routes/                # Route definitions
├── docs/                  # Swagger generated docs
└── main.go                # Entry point

frontend/
├── src/
│   ├── pages/             # Route pages
│   ├── types/             # TypeScript interfaces
│   ├── api.ts             # Axios client
│   └── App.tsx            # Root component
└── package.json
```

## Database Schema

| Table                  | Description                                   |
|------------------------|-----------------------------------------------|
| `users`                | Authentication (email + password hash)        |
| `organizations`        | Top-level tenant grouping                     |
| `organization_members` | Many-to-many users ↔ orgs with roles          |
| `projects`             | Projects scoped within an organization        |
| `test_suites`          | Hierarchical folder-like organization         |
| `test_cases`           | Core entity (steps as JSONB, tags, status)    |
| `test_case_history`    | Immutable audit log for test case changes     |
| `test_plans`           | Named groupings of test runs                  |
| `test_runs`            | Execution instances linked to a plan          |
| `test_run_cases`       | Junction table for runs ↔ cases               |
| `test_results`         | Individual execution results per case         |
| `result_attachments`   | File attachments linked to a test result      |
| `defects`              | Defect/bug tracking linked to results         |
