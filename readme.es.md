# Golang Agnostic Template

![Go](https://img.shields.io/badge/Go-1.24-blue)
![License](https://img.shields.io/badge/License-MIT-redi)
![Database](https://img.shields.io/badge/Database-SurrealDB-purple)

**Golang Agnostic Template** es una plantilla de backend en Go diseñada para construir aplicaciones modernas, escalables y mantenibles utilizando la **arquitectura hexagonal** (Puertos y Adaptadores). Este proyecto proporciona una base sólida para desarrollar aplicaciones web con un enfoque en la limpieza del código, la modularidad y las mejores prácticas de desarrollo.

## Características principales

- **Arquitectura Hexagonal**: Separa la lógica de negocio (dominio) de la infraestructura, facilitando la escalabilidad, el mantenimiento y el reemplazo de componentes externos (como bases de datos o frameworks web).
- **Patrones de diseño**:
  - **Repository Pattern**: Abstrae el acceso a datos mediante interfaces definidas en `src/application/domain/repository`.
  - **Dependency Injection**: Inyecta dependencias en los componentes, como handlers y servicios, para mejorar la testabilidad.
  - **Factory Pattern**: Centraliza la creación de entidades del dominio en `src/application/domain/factory.go`.
  - **DTO (Data Transfer Objects)**: Utiliza DTOs (`src/application/domain/dto/register_user.go`) para manejar datos de entrada/salida de forma segura y validada.
- **Modularidad**: Estructura clara con separación de dominio (`src/application/domain`), adaptadores (`src/application/actors`), e infraestructura (`src/pkg`).
- **Contenerización**: Incluye soporte para Docker mediante `docker-compose.yaml`, facilitando el despliegue en entornos locales y de producción.

## Componentes tecnológicos
- [X] **Gin**: Framework web.
- [X] **SurrealDB Singleton**: Base de datos moderna multimodal.
- [X] **Zappier**: Framework para manejo de Logs.
- [ ] **NATs**: Framework para sistema de mensajería nativo de la nube.
- [ ] **SurrealDB Multitenant**: Manejo de Namespaces y escrituras.
- [ ] **RxGo**: Reactive programming.

## Estructura del proyecto

### Project Structure

```
### Project Structure

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

## Instalación

1. **Clonar el repositorio**:
   ```bash
   git clone https://github.com/dany0814/golang-agnostic-template.git &&
   cd golang-agnostic-template
   ```

2. **Configurar el entorno**:
   
   Asegúrate de tener Go 1.24+ instalado e Instala las dependencias:
      ```bash
      go mod tidy
      ```
   Crea el archivo .env en el directorio raíz (Guíate del archivo /.env.example):
      ```bash
      cp -p .env.example .env
      ```

3. **Ejecutar SurrealDB con Docker**:
   
   Usa `docker-compose` para iniciar los servicios:
     ```bash
      docker-compose up -d --build
      ```

4. **Ejecutar localmente**:
   
   Compila la aplicación y ejecuta el servidor desde el directorio raíz:
     ```bash
      go run main.go
      ```

## Uso

- **Aplicaciones Web Escalables**: La aplicación expone endpoints definidos en `src/application/actors/web/router.go`. Por ejemplo, puedes registrar un usuario enviando una solicitud POST a `/users/register` con un cuerpo JSON basado en `register_user.go`.
- **Base de datos**: SurrealDB almacena los datos de los usuarios, con entidades definidas en `src/application/domain/model/user.go`.

## Licencia

Este proyecto está licenciado bajo la [Licencia MIT](LICENSE).