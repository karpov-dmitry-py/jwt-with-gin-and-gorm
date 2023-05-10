## JWT with Gin and GORM

The app runs a simple HTTP server with basic JWT authorization implementation

### app commands:

- `make run` - run app and db 
- `make stop` - stop app and db
- `make kill` - delete app and db containers and volumes
- `make log` - show app streaming logs
- `make install-linter` - install golangci-lint binary into the project
- `make lint` - run linter
- `make lint-fast` - run linter with --fast option

### app routes:

- GET `/alive` - app's health check
- POST `/signup` - create a user
- todo - login