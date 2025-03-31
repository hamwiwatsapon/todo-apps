# Todo Projects

This repository contains a full-stack Todo application with a **backend** built using Go (Fiber framework) and a **frontend** built using Next.js.

## Project Structure

````plaintext
todo-projects/
├── backend/         # Backend service
│   ├── cmd/        # Main entry point for the backend
│   │   └── main.go # Main application file
│   ├── docs/       # Swagger API documentation
│   ├── internal/   # Internal application logic
│   │   ├── delivery/   # HTTP handlers
│   │   ├── domain/     # Domain models and interfaces
│   │   ├── repository/ # Data access layer
│   │   └── usecase/    # Business logic layer
│   ├── package/    # Shared packages
│   └── database/   # Database connection and migrations
├── frontend/        # Frontend service
│   └── todo-apps/   # Next.js application
│       ├── public/      # Static assets
│       ├── src/         # Source code
│       │   ├── app/        # Application entry point
│       │   ├── components/ # Reusable UI components
│       │   ├── feature/    # Feature-specific components and logic
│       │   └── lib/        # Utility functions
│       ├── .env.example    # Example environment variables
│       ├── package.json    # Frontend dependencies and scripts
│       └── tsconfig.json   # TypeScript configuration
└── README.md         # Project documentation

## Backend

The backend is built using Go and the Fiber framework. It provides a RESTful API for managing todos.

### Backend Features

- CRUD operations for todos
- SQLite database integration using GORM
- Swagger API documentation

### Setup

1. Navigate to the `backend` directory:

   ```bash
   cd backend
   go mod tidy
   go run cmd/main.go
````

### Directory Overview

cmd/main.go: Entry point for the backend application.
internal/domain: Contains domain models and interfaces.
internal/repository: Implements data access logic using GORM.
internal/usecase: Contains business logic.
internal/delivery/http: HTTP handlers for the API.
package/database: Database connection and migrations.

## Frontend

The frontend is a Next.js application that provides a user interface for managing todos.

### Frontend Features

Add, update, delete, and view todos
Responsive design
Integration with the backend API

### Setup

1. Navigate to the frontend/todo-apps directory:
   cd frontend/todo-apps

2. Install dependencies:
   npm install

3. Run the development server:
   npm run dev

4. Open the application in your browser at [http://localhost:3000](http://localhost:3000).

### Frontend Directory Overview

src/app: Application entry point and layout.
src/components: Reusable UI components.
src/feature/todos: Todo-specific components and API logic.
src/lib: Utility functions.

## Environment Variables

### Backend ENV

Create a .env file in the backend directory with the following variables:
DATABASE_URL=sqlite://todo.db

### Frontend ENV

Copy .env.example to .env in the frontend/todo-apps directory and update the BACKEND_URL:
BACKEND_URL=http://localhost:4444

### License

This project is licensed under the Apache 2.0 License.
You can replace or modify the content as needed to better suit your project.
