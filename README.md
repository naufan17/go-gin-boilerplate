# Boilerplate Golang REST API
Minimalist project structure using Gin to build REST API

## Table of Content
- [Preparation](#preparation)
- [Run Server Locally](#run-server-locally)
- [Run Server with Docker Compose](#run-server-with-docker-compose)
- [Feature](#feature)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Default API Endpoint](#default-api-endpoint)
- [Response Format](#response-format)

## Preparation
- Clone this repository
```
git clone https://github.com/naufan17/go-gin-boilerplate.git
```
- Enter directory project
```
cd go-gin-boilerplate
```
- Delete git
```
rm -rf .git
```
- Rename .env.example to .env and fill in the section that must be filled
- Custom your project name
- Install dependencies
```
go mod download
```
- Run database migration
```
go run cmd/database/migration/main.go
```
- Run database seeder
```
go run cmd/database/seeder/main.go
```

## Run Server Locally
- Run server
```
go run cmd/server/main.go
```

## Run Server with Docker Compose
- Run server
```
docker compose up --build
```
- Stop server
```
docker compose down
```

## Feature
- **Database**: using https://gorm.io/driver/postgres
- **ORM**: using https://gorm.io/
- **Authentication**: using https://github.com/dgrijalva/jwt-go
- **Validation**: using https://github.com/go-playground/validator/v10
- **Rate Limit**: using https://golang.org/x/time/rate
- **Security**: using https://github.com/gin-contrib/secure or https://github.com/danielkov/gin-helmet
- **CORS**: using https://github.com/gin-contrib/cors
- **Compression**: using https://github.com/gin-contrib/gzip
- **Environtment variables**: using https://github.com/joho/godotenv
- **API documentation**: using https://github.com/swaggo/swag, https://github.com/swaggo/gin-swagger and https://github.com/swaggo/files
- **Dependency management**: using https://github.com/golang/dep

## Project Structure
```
cmd\
  |--server\       # Command to run server
  |--database\     # Command database migration and seeder          
config\            # Configuration
database\
  |--seeders\      # Database seeder
internal\
  |--dtos\         # Data transfer object
  |--handlers\     # Request handlers
  |--middlewares\  # Middleware
  |--models\       # Database models
  |--repositories\ # Database queries
  |--services\     # Business logic
pkg\
  |--auth\         # Authentication
  |--util\         # Utility function
route\             # API routes
```

## API Documentation
To view the API documentation, open the following link:
<br/>
``GET /swagger/index.html`` - View API documentation

## Default API Endpoint
**Auth routes**:
<br/>
``POST /api/auth/register`` - Create new account
<br/>
``POST /api/auth/login`` - Login to existing account
<br/>
``GET /api/auth/refresh`` - Refresh access token
<br/>
``GET /api/auth/logout`` - Logout account

**Account routes**:
<br/>
``GET /api/account/profile`` - Get current account profile
<br/>
``GET /api/account/session`` - Get current account session
<br/>
``POST /api/account/update-profile`` - Update current account profile
<br/>
``POST /api/account/update-password`` - Update current account password

## Response Format
- **Success Response**:
```
{
  "data": {
    "key": "value",
  }
}
```
OR
```
{
  "message": "success message",
}
```
- **Error Response**:
```
{
  "error": "error message",
}
```