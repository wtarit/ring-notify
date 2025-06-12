# Ring Notify Backend

## Running Docker Image

`docker container run --rm --env-file .env -p 1323:1323 ghcr.io/wtarit/ring-notify-backend:0.0.1`

## Build Docker Image

Inside `api/` directory
`docker build --tag ring-notify-backend .`
