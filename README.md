# ps-tag-onboarding-go

# How to run
At the root folder of the project run `docker compose up --build` to run the mongodb and the go application.

Once you see these messages in the console, the application is ready to be used:
```
ps-tag-onboarding-go-api-1             | Starting Application!
ps-tag-onboarding-go-api-1             | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
ps-tag-onboarding-go-api-1             | 
ps-tag-onboarding-go-api-1             | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
ps-tag-onboarding-go-api-1             |  - using env:  export GIN_MODE=release
ps-tag-onboarding-go-api-1             |  - using code: gin.SetMode(gin.ReleaseMode)
ps-tag-onboarding-go-api-1             | 
ps-tag-onboarding-go-api-1             | [GIN-debug] GET    /users/:id                --> github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller.UserController.GetUserById-fm (3 handlers)ps-tag-onboarding-go-api-1             | [GIN-debug] POST   /users                    --> github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller.UserController.CreateUser-fm (3 handlers) 
ps-tag-onboarding-go-api-1             | [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
ps-tag-onboarding-go-api-1             | Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
ps-tag-onboarding-go-api-1             | [GIN-debug] Listening and serving HTTP on :8080
```

# API Endpoints

## Create a new user
```
curl -X POST http://localhost:8080/users -H 'Content-Type: application/json' -d '{"first_name":"John","last_name":"Doe","email":"JohnDoe@test.com","age":24}'
```

## Get User By id
```
curl -X GET http://localhost:8080/users/{id}
```
NOTE: Replace `{id}` with the guid returned when creating the user

# Technical proposal

## Gin Web Framework

### What
Gin Web Framework is known for its high performance and HTTP requests processing speed. The framework is sufficient for building a simple application with simple endpoints for CRUD operations.

### Why
For this project, Gin has features that can help with:
- Converting the JSON playload to a struct and vice versa
- Routing API endpoints to the appropriate handler functions
- Middleware support, currently being used to return the status code and error message. This can be extended to include other functionalities like logging and authentication.

## Testify Framework

### What
Testify framework allows for easy testing of the application. It provides a set of helper functions that can be used to test the application. The framework has a number of packages that can be used to test the application including:
- Assert package: This package is used to validate the conditions or test the behaviour of the code
- Mock package: This package is used to mock the database and other dependencies

### Why
The benefits of using testify to test the application include:
- Easy to use UI for debugging
- Assertion are easy to read and understand
- Mocking the database and other dependencies when testing

Testify also has a huge community of contributors and users, so there are many documentations and resources online to help with any issues that may arise. It is also used by many TAG services, so I can get some experience with using this framework for testing.

## MongoDB

### What
MongoDB is a document database known for its flexibility and scalability. It is a good choice for this application as it allows for easy storage and retrieval of data.

### Why
- Storing data in JSON format, which is easy to work with
- Free to use

MongoDB is used by many services in TAG, which I haven't worked with before. So this is a good opportunity to get familiar with it.