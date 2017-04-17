# Auth REST API
A RESTful API providing authentication through JWT

It is written in Go, mainly using **labstack/echo** and **gorm**

## Installation & Run
```bash
# Download this project
go get github.com/m5r/auth-rest-api
```

```bash
# Build and Run
cd auth-rest-api
go build
./auth-rest-api

# API Endpoint : http://127.0.0.1:3000
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler           // Our API core handlers
│   │   ├── auth.go       // /auth route handler
│   │   ├── common.go     // Common response functions
│   │   ├── index.go      // / route handler
│   │   ├── signup.go     // /signup route handler
│   │   ├── refresh.go    // /refresh route handler
│   │   └── restricted.go // /restricted route handler
│   └── model
│       └── model.go      // Models for our application
├── config
│   └── config.go         // Configuration
└── main.go
```

## API

#### /
* `GET` : Get all projects

#### /auth
* `POST` : Log in

#### /signup
* `POST` : Sign up

#### /refresh
* `POST` : Refresh the JWT

#### /restricted
* `GET` : Get the restricted data
## Todo

- [ ] Document the routes.
- [ ] Document the configuration.
- [ ] Write the tests for all APIs.
