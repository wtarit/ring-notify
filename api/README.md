# Ring Notify Backend

A Go-based API server for sending Firebase Cloud Messaging (FCM) notifications to trigger calls.

## Features

- **User Management**: Create users with FCM tokens and get API keys
- **FCM Notifications**: Send push notifications to trigger calls

### Accessing Swagger UI

When the server is running, you can access the interactive Swagger UI at:

```
http://localhost:1323/swagger/index.html
```

### API Endpoints

- `GET /` - Health check endpoint
- `POST /user/create` - Create a new user with FCM token
- `POST /notify/call` - Send FCM notification (requires Bearer token)

### Regenerating Documentation

To regenerate the Swagger documentation after making changes to the API:

```bash
./generate-docs.sh
```

## Running the Application

### Local Development

1. Install dependencies:

```bash
go mod download
```

2. Set up environment variables (copy `.env.sample` to `.env` and configure)

3. Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:1323`

### Docker

#### Running Docker Image

`docker container run --rm --env-file .env -p 1323:1323 ghcr.io/wtarit/ring-notify-backend:0.0.1`

#### Build Docker Image

Inside `api/` directory:

```bash
docker build --tag ring-notify-backend .
```
