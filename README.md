# Blogging Platform API
A RESTful API for a blogging platform .

## Prerequisites
- Go 1.21.11 or higher
- Docker and Docker Compose

## Environment Setup

1. Clone the repository
2. Create a `.env` file in the root directory with the following variables:

```env
PORT=YOUR_SPECIFIC_PORT_FOR_WEB_SERVER
MONGO_URI=YOUR_MONGO_DB_CONNECTION_STRING
MONGO_DB_NAME=YOUR_MONGODB_NAME
```

## Running the Application

### Using Docker (Recommended)

1. Start the MongoDB container:
```bash
docker compose up -d
```

2. Run the application:
```bash
go run cmd/server/main.go
```

### Manual Setup

1. Install dependencies:
```bash
go mod download
```

2. Ensure MongoDB is running and accessible

3. Run the application:
```bash
go run cmd/server/main.go
```


## Development
The project uses the following main dependencies:
- Gin Web Framework
- MongoDB Go Driver
- Godotenv for environment variables