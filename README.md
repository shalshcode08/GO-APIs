# Go-APIs

A RESTful API built with Go for managing users, products, and orders. This project provides a complete e-commerce backend with authentication, product management, and order processing capabilities.

## Features

- **User Authentication**: JWT-based authentication with secure password hashing
- **User Management**: User registration and login endpoints
- **Product Management**: Create and retrieve products
- **Order Management**: Create orders and order items
- **Cart/Checkout**: Shopping cart checkout functionality
- **Database Migrations**: Version-controlled database schema management
- **Input Validation**: Request payload validation using go-playground/validator

## Tech Stack

- **Language**: Go 1.24+
- **Web Framework**: Gorilla Mux
- **Database**: MySQL
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt (golang.org/x/crypto)
- **Validation**: go-playground/validator
- **Migrations**: golang-migrate/migrate
- **Configuration**: godotenv

## Prerequisites

- Go 1.24 or higher
- MySQL 5.7+ or MySQL 8.0+
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/shalshcode08/GO-APIs.git
cd Go-APIs
```

2. Install dependencies:
```bash
go mod download
```

3. Install golang-migrate (if not already installed):
```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
# Database Configuration
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=GOAPI

# Server Configuration
PORT=8080
PUBLIC_HOST=http://localhost

# JWT Configuration
JWT_SECRET=your-secret-key-change-this-in-production
JWT_EXP=604800  # 7 days in seconds (3600*24*7)
```

**Note**: Make sure to change `JWT_SECRET` to a secure random string in production environments.

## Database Setup

1. Create a MySQL database:
```sql
CREATE DATABASE GOAPI;
```

2. Run database migrations:
```bash
# Run all migrations
make migrate-up

# Or manually:
go run cmd/migrate/main.go up
```

3. To rollback migrations:
```bash
make migrate-down

# Or manually:
go run cmd/migrate/main.go down
```

## Running the Project

### Using Makefile (Recommended)

```bash
# Build and run
make run

# Or separately:
make build
./bin/GO-APIs
```

### Manual Run

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080` (or the port specified in your `.env` file).

## API Endpoints

All endpoints are prefixed with `/api/v1`.

## Testing

Run all tests:
```bash
make test

# Or manually:
go test -v ./...
```

## Development

### Creating New Migrations

```bash
make migration create_migration_name

# Example:
make migration add_cart_table
```

This will create both `.up.sql` and `.down.sql` files in `cmd/migrate/migrations/`.

### Building

```bash
make build
```

This creates a binary in the `bin/` directory.

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_USER` | MySQL username | `` |
| `DB_PASSWORD` | MySQL password | `` |
| `DB_HOST` | MySQL host | `` |
| `DB_PORT` | MySQL port | `` |
| `DB_NAME` | Database name | `` |
| `PORT` | Server port | `` |
| `PUBLIC_HOST` | Public host URL | `` |
| `JWT_SECRET` | JWT signing secret | `` |
| `JWT_EXP` | JWT expiration in seconds | ``|