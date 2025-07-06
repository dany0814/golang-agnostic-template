# Golang Agnostic Template

![Go](https://img.shields.io/badge/Go-1.24-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Database](https://img.shields.io/badge/Database-SurrealDB-purple)

**Golang Agnostic Template** is a Go backend template designed for building modern, scalable, and maintainable applications using **hexagonal architecture** (Ports and Adapters). This project provides a solid foundation to develop web applications focused on clean code, modularity, and best development practices.

## Key Features

- **Hexagonal Architecture**: Separates business logic (domain) from infrastructure, enabling scalability, maintainability, and easy replacement of external components (e.g., databases or web frameworks).
- **Design Patterns**:
  - **Repository Pattern**: Abstracts data access through interfaces defined in `src/application/domain/repository`.
  - **Dependency Injection**: Injects dependencies into components, such as handlers and services, to improve testability.
  - **Factory Pattern**: Centralizes the creation of domain entities in `src/application/domain/factory.go`.
  - **DTO (Data Transfer Objects)**: Uses DTOs (`src/application/domain/dto/register_user.go`) to securely handle input/output data with validation.
- **Modularity**: Clear structure with separation of domain (`src/application/domain`), adapters (`src/application/actors`), and infrastructure (`src/pkg`).
- **Containerization**: Includes Docker support via `docker-compose.yaml`, simplifying deployment in local and production environments.

## Technological Components
- [x] **Gin**: Web framework for building RESTful APIs.
- [x] **SurrealDB Singleton**: Modern multimodal database.
- [x] **Zap**: Framework for structured logging.
- [ ] **NATS**: Cloud-native messaging system framework.
- [ ] **SurrealDB Multitenant**: Namespace and write management.
- [ ] **RxGo**: Reactive programming library.

## Project Structure

The project follows a modular structure aligned with hexagonal architecture principles, as shown below:

```
- 📁 **golang-agnostic-template/**
  - 📄 `docker-compose.yaml` - Docker Compose configuration
  - 📁 **src/**
    - 📁 **application/**
      - 📁 **actors/**
        - 📁 **db/**
          - 📄 `user_repository.go` - User repository implementation
        - 📁 **web/**
          - 📄 `handler.go` - API handlers
          - 📄 `router.go` - Gin router configuration
      - 📁 **domain/**
        - 📁 **business/**
          - 📄 `user.go` - User business logic
        - 📁 **dto/**
          - 📄 `register_user.go` - User registration DTO
        - 📁 **model/**
          - 📄 `user.go` - User entity
        - 📁 **repository/**
          - 📄 `user_repository.go` - User repository interface
        - 📁 **service/**
          - 📄 `user.go` - User service logic
          - 📄 `organization.go` - Organization service logic
        - 📁 **utils/**
          - 📄 `constants.go` - Domain constants
          - 📄 `errors.go` - Custom error handling
          - 📄 `utils.go` - Utility functions
    - 📁 **pkg/**
      - 📁 **database/**
        - 📄 `surrealdb.go` - SurrealDB adapter
      - 📁 **webserver/**
        - 📄 `server.go` - Gin web server configuration
```

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/dany0814/golang-agnostic-template.git &&
   cd golang-agnostic-template
   ```

2. **Set up the environment**:
   
   Ensure Go 1.24+ is installed and Install dependencies:
      ```bash
      go mod tidy
      ```
   
   Create the `.env` file in the root directory (use `.env.example` as a guide):
      ```bash
      cp -p .env.example .env
      ```

3. **Run SurrealDB with Docker**:
   
   Use `docker-compose` to start the services:
     ```bash
     docker-compose up -d --build
     ```

4. **Run locally**:
   
   Compile application and run the server from root directory:
     ```bash
     go run main.go
     ```

## Usage

- **Scalable Web Applications**: The application exposes endpoints defined in `src/application/actors/web/router.go`. For example, you can register a user by sending a POST request to `/users/register` with a JSON body based on `register_user.go`.
- **Database**: SurrealDB stores user data, with entities defined in `src/application/domain/model/user.go`.

## License

This project is licensed under the [MIT License](LICENSE).