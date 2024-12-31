# FAMPAY-BACKEND-ASSIGNMENT

This project is designed to fetch data from the YouTube API v3 and display it in paginated results. It uses a microservices architecture for scalability and ease of management.

---

## Services

The services used in this project include:

### 1. **App**
- **Purpose**: The main API service, written in Golang.
- **Features**:
  - Exposes REST API endpoints for video fetching.
  - Triggers Temporal workflows.
  - Connects to the PostgreSQL database.
- **Ports**: Exposed at `8080`.

### 2. **Postgres**
- **Purpose**: The primary database for storing video data.
- **Features**:
  - Provides indexing and relational data handling.
  - Ensures high availability and performance.
- **Ports**: Exposed at `5432`.

### 3. **Temporal**
- **Purpose**: Workflow orchestration.
- **Features**:
  - Handles scheduling and execution of background jobs.
  - Facilitates reliable task retries and error handling.
- **Ports**: Exposed at `7233`

### 4. **Temporal Admin Tools**
- **Purpose**: Administrative tools for managing Temporal workflows.

### 5. **Temporal Web**
- **Purpose**: Web UI for monitoring Temporal workflows.
- **Ports**: Exposed at `8088`.

---

## Getting Started

### Prerequisites
- **Required**:
  - Docker and Docker Compose
  - Git

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/NayanPahuja/fampay-task.git
   ```

2. Navigate to the project directory:
   ```bash
   cd fampay-task
   ```
3. Edit the `.env.example` to your liking if running it locally or add the environment variables to the docker-compose file:

4. Build and start the services:
   ```bash
   docker-compose up --build 
   ```

---

## Usage

### Services Access

- **Web API Service**: [http://localhost:8080/health](http://localhost:8080)
  - Navigate to this URL to access the API service and check it's health.

- **Swagger UI**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
  - Use this for testing the API endpoints.

- **Temporal Web**: [http://localhost:8088](http://localhost:8088)
  - Access this for monitoring workflows.

---

## API Endpoints

### Video Endpoints

1. **GET /api/v1/videos**
   - Fetches all videos from the database with support for pagination using limit and offset.

2. **GET /api/v1/videosv2**
   - Fetches videos from the database with cursor-based pagination for scalability.

---

## Directory Structure

```plaintext
📦 fampay-task
├─ .env.example
├─ .gitignore
├─ Dockerfile
├─ Makefile
├─ README.md
├─ cmd
│  ├─ docs
│  │  ├─ docs.go
│  │  ├─ swagger.json
│  │  └─ swagger.yaml
│  └─ main.go
├─ config
│  └─ config.go
├─ db
│  └─ db.go
├─ docker-compose.yml
├─ docs
│  ├─ docs.go
│  ├─ swagger.json
│  └─ swagger.yaml
├─ entrypoint.sh
├─ go.mod
├─ go.sum
├─ internal
│  ├─ handlers
│  │  ├─ health_handler.go
│  │  └─ video_handler.go
│  ├─ models
│  │  └─ videoModel.go
│  ├─ repositories
│  │  └─ video_repo.go
│  ├─ routes
│  │  └─ routes.go
│  ├─ services
│  │  └─ video_service.go
│  ├─ temporal
│  │  ├─ client.go
│  │  └─ worker.go
│  ├─ utils
│  │  ├─ cursor.go
│  │  └─ youtube.go
│  └─ workflows
│     ├─ yt-activity.go
│     └─ yt-workflow.go
└─ migrate
   ├─ main.go
   └─ migrations
      ├─ create-youtube-table.up.sql
      ├─ create-youtube-table.down.sql
      └─ ...
```

---

## Milestones

-  Create a worker that fetches the latest videos from the YouTube API.
-  Schedule Temporal workflows that periodically trigger activities.
-  Design models with GORM for storing video data.
-  Setup indexing for database tables.
-  Implement API endpoints for fetching videos.
-  Add pagination using limit and offset.
-  Enable cursor-based pagination for scalability.
-  Integrate Swagger for API testing.
-  Use Docker Compose to manage multi-service architecture.
---

## Future Improvements
- Create a basic frontend that gets results from the backend
- Right now insertion of videos is not cached (if we restart the server, same videos are tried to be inserted) ->In progress

## Reasons for Service Selection

### Why PostgreSQL?
- **Benefits**:
  - Supports advanced indexing for efficient querying.
  - Scalable and well-suited for relational data models.
  - Easy to manage

### Why Temporal?
- **Advantages**:
  - Simplifies orchestration of complex workflows.
  - Ensures reliable retries and error handling.
  - Offers a robust web UI for monitoring workflows.

