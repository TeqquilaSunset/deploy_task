# Gin REST API Backend Starter

A robust "kickstarter" template for building backend applications using the Gin web framework in Go. This project provides a well-organized structure with essential features like MongoDB integration, JWT-based authentication, and middleware support. It aims to simplify the setup of a functional backend environment, allowing developers to focus on application-specific functionality. 

While it provides a comprehensive starting point, it remains adaptable for further customization and expansion, allowing developers to incorporate additional architectural patterns or advanced features as needed. This makes it an effective foundation for creating scalable and maintainable web services in Go.

## Project Structure
```bash
gin-backend-starter/
├── assets      # Static files and resources
├── configs     # Configuration files and setup, these would execute during the pre-start process of the server
├── helpers     # Utility functions and common helpers
├── infra       # Infrastructure-related code (database, logger, etc. that would be used throughout the app)
├── models      # Data models and database schemas
├── routers     # HTTP route definitions and handlers mapping
├── services    # Business logic and service layer implementations
├── .env        # Application config
├── .env.example
├── .gitignore
├── air.toml
├── go.mod
├── go.sum
├── LICENSE
├── main.go     # Main entry point for the application
└── README.md
```

## Features

- Structured project layout for scalability
- Configuration management
- MongoDB integration
- JWT authentication
- User management
- Middleware support
- Logging
- Environment variable configuration

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/imshawan/gin-backend-starter.git
    cd gin-backend-starter
    ```

2. **Copy the example environment file and update the values:**

    ```bash
    cp .env.example .env
    ```

3. **Install dependencies:**

    ```bash
    go mod tidy
    ```

4. **Run the application:**

    ```bash
    go run main.go
    ```

## Environment Variables

Update the `.env` file with your specific configurations before starting the app


## Starter API Routes

- `POST /api/users/register`: Register a new user
- `POST /api/auth/sign-in`: User login
- `GET /api/users`: Get the currently logged-in user (protected route)
- `GET /health`: Check the server status, if running

## Major Dependencies

- [Gin Web Framework](https://github.com/gin-gonic/gin) for creating the REST APIs
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) for communicating with MongoDB NoSql Database
- [JWT Go](https://github.com/golang-jwt/jwt/) for authentication purposes
- [Viper](github.com/spf13/viper) for config management throughout the application

## Development

This project uses [Air](https://github.com/air-verse/air) for live reloading during development. To use it:

1. **Install Air:**

    ```bash
    go install github.com/air-verse/air@latest
    ```

2. **Run the project using Air:**

    ```bash
    air
    ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b improvement/AmazingImprovements`)
3. Commit your changes (`git commit -m 'Added some amazing improvements'`)
4. Push to the branch (`git push origin improvement/AmazingImprovements`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
