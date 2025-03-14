# Golang RestAPI with Case CoffeeShop

A simple, yet powerful Go web application for managing a coffee shop menu with RESTful endpoints. This application demonstrates the use of Go, Gin framework, configuration management, logging, and containerization.

## Features

- RESTful API for coffee shop menu management
- Web interface for displaying coffee items
- Configuration-based coffee menu
- Structured logging using Zap
- Docker support for easy deployment

## Tech Stack

- Go (Golang)
- Gin Web Framework
- Viper (Configuration management)
- Zap (Structured logging)
- Docker & Docker Compose

## Project Structure

```
.
├── .env                  # Environment variables
├── .gitignore            # Git ignore file
├── Dockerfile            # Docker build instructions
├── README.md             # This file
├── coffee/               # Coffee package
│   ├── coffee.go         # Coffee data models and functions
│   └── coffee_test.go    # Tests for coffee package
├── config.json           # Configuration file containing coffee menu
├── docker-compose.yml    # Docker compose configuration
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── logs/                 # Log directory
│   └── logs.txt          # Application logs
├── main.go               # Main application entry point
├── main_test.go          # Tests for main package
└── templates/            # HTML templates
    └── index.html        # Coffee menu display template
```

## API Endpoints

- `GET /ping`: Health check endpoint that returns a welcome message
- `GET /home`: Renders an HTML page displaying the coffee menu
- `GET /coffee`: Returns the coffee menu data

## Configuration

The application uses a `config.json` file to store the coffee menu items:

```json
{
     "coffeelist": [
          {
               "name": "Latte",
               "price": 3
          },
          {
               "name": "Cappuccino",
               "price": 2.5
          },
          {
               "name": "Flat White",
               "price": 2.25
          }
     ]
}
```

You can modify this file to add, remove, or update coffee items.

## Environment Variables

The application uses the following environment variables defined in `.env`:

- `APP_PORT`: The port on which the application will run (default: 8085)

## Running the Application

### Locally

1. Ensure you have Go installed (version 1.16+)
2. Clone the repository
3. Navigate to the project directory
4. Run the application:

```bash
go mod tidy
go run main.go
```

5. Access the application at http://localhost:8085

### Using Docker

1. Ensure you have Docker and Docker Compose installed
2. Navigate to the project directory
3. Build and start the containers:

```bash
docker-compose up -d
```

4. Access the application at http://localhost:8085

## Logging

Application logs are written to `logs/logs.txt`. The logging system uses Zap for structured logging and captures important events and errors within the application.

## Testing

To run the tests:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.