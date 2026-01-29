# Sandwich
A project management app for personal use.

## Features

- Manage projects and categories.
- Provides a REST API to interact with the application.
- A CLI to manage the projects and categories from the terminal.

## API Endpoints

### Categories

- `GET /api/category`: Get all categories.
- `POST /api/category`: Add a new category.
- `PATCH /api/category/{id}`: Update a category.
- `DELETE /api/category/{id}`: Delete a category.
- `GET /api/category/{id}`: Get all projects for a category.

### Projects

- `GET /api/category/{catId}/projects/{id}`: Get a specific project.
- `POST /api/category/{catId}/projects`: Add a new project to a category.
- `PATCH /api/category/{catId}/projects/{id}`: Update a project.
- `DELETE /api/category/{catId}/projects/{id}`: Delete a project.
- `GET /api/projects/random`: Get a random project.
- `GET /api/projects/random/{num}`: Get a specified number of random projects.

## CLI Usage

The CLI is built using `cobra`.

### Root Command

```bash
sandwich [command]
```

### Subcommands

The available subcommands are:
- `category`: Manage categories.
- `project`: Manage projects.
- `web`: Start the web server.

#### Category Command

```bash
sandwich category [subcommand]
```

- `add`: Add a new category.
- `list`: List all categories.
- `update`: Update a category.
- `delete`: Delete a category.

#### Project Command

```bash
sandwich project [subcommand]
```

- `add`: Add a new project.
- `list`: List all projects.
- `update`: Update a project.
- `delete`: Delete a project.
- `random`: Get a random project.

## Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/srynprjl/sandwich.git
    ```
2.  Go to the project directory:
    ```bash
    cd sandwich
    ```
3.  Build the project:
    ```bash
    go build
    ```
4.  Run the application:
    ```bash
    ./sandwich
    ```
