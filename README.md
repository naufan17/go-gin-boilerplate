# Bolierplate Golang REST API
Minimalist project structure using Gin to build REST API

## Table of Content
- [Preparation](#preparation)
- [Run Server](#run-server)
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

## Run Server
- Run server
```
go run cmd/api/main.go
```

## Feature
- **ORM**: using https://gorm.io/
- **Authentication**: using https://github.com/dgrijalva/jwt-go
- **Rate Limit**: using https://golang.org/x/time/rate
- **Security**: using https://github.com/gin-contrib/secure
- **CORS**: using https://github.com/gin-contrib/cors
- **Environtment variables**: using https://github.com/joho/godotenv
- **API documentation**: using https://github.com/swaggo/swag
- **Dependency management**: using https://github.com/golang/dep

## Project Structure
```
cmd\               # Commands
database\          # Database seed
internal\          # Source code
  |--configs\      # Configuration
  |--controllers\  # API controllers
  |--dtos\         # Data transfer object
  |--middlewares\  # Middleware
  |--models\       # Database models
  |--repositories\ # Database queries
  |--routes\       # API routes
  |--services\     # Business logic
  |--utils\        # Utility function
```

## API Documentation
To view the API documentation, open the following link:
<br/>
``GET /swagger/index.html`` - View API documentation

## Default API Endpoint
**Auth routes**:
<br/>
``POST /api/v1/auth/register`` - Create new account
<br/>
``POST /api/v1/auth/login`` - Login to existing account

**Account routes**:
<br/>
``GET /api/v1/account/profile`` - Get current account profile
<br/>

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