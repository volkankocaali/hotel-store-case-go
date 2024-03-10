# Hotel Store Case

## Description

This application is an application that provides services such as creating, updating, listing reservations that contain authentication made with go.

## Installation
1. Clone the repository
2. Run `docker-compose up -d` command
3. Run `cp .env.dist .env`
4. Update the .env file with your database settings
5. Run `go run cmd/api/server.go` or `go build cmd/api/server.go` to start the server

### Postman Collection

[App.postman_collection.json](App.postman_collection.json)

## Configuration

Database settings can be configured in the .env file.

## License

This project is licensed under the MIT license.
