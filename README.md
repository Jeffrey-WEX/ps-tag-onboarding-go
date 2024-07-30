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
ps-tag-onboarding-go-api-1             | [GIN-debug] GET    /v1/users/:id                --> github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller.UserController.GetUserById-fm (3 handlers)ps-tag-onboarding-go-api-1             | [GIN-debug] POST   /v1/users                    --> github.com/Jeffrey-WEX/ps-tag-onboarding-go/internal/controller.UserController.CreateUser-fm (3 handlers) 
ps-tag-onboarding-go-api-1             | [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
ps-tag-onboarding-go-api-1             | Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
ps-tag-onboarding-go-api-1             | [GIN-debug] Listening and serving HTTP on :8080
```

# API Endpoints

## Create a new user
```
curl -X POST http://localhost:8080/v1/users -H 'Content-Type: application/json' -d '{"first_name":"John","last_name":"Doe","email":"JohnDoe@test.com","age":24}'
```

## Get User By id
```
curl -X GET http://localhost:8080/v1/users/{id}
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

### Comparison with other framework
I have also considered using Echo framework. Both framework are maintained and updated frequently and the community for both framework are very active so it is easy to look for help if needed, but Gin community is overall bigger so there would be more resources online for Gin. Echo's dynamic routing allows for more flexibility and customization, while Gin provides a simpler and straightforward static routing. The reason I chose Gin over Echo is because:
- This project have simple routing requirements so Gin's static routing is sufficient and offer better performance
- Gin has a smaller memory footprint due to the static routing
- Gin uses a middleware handler chain, allowing for more control over the order of executions
- Mainly to get familiarity with Gin as it is used by many TAG services

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

### Comparison with other framework
I have also considered using Go's built-in testing package. The built-in Go testing package is sufficient for testing the application, but it lacks some features that are available in Testify. Testify provides a more user-friendly interface for testing and has more features that can be used to test the application. Specifically, testify has a more comprehensive set of assertion functions that can be used to test the application. Testify also allows assertion to include messages, making the test easier to understand. That is why I chose Testify over the built-in Go testing package.

## MongoDB

### What
MongoDB is a document database known for its flexibility and scalability. It is a good choice for this application as it allows for easy storage and retrieval of data.

### Why
- Storing data in JSON format, which is easy to work with
- Free to use

### Comparison with other database options
The other option I have looked for is using an SQL database like postgresql. The main reason I chose MongoDB over postgresql are because:
- The data is stored in JSON format, which is to integrate with the service, as the data being sent in is also in JSON format
- The project only require simple queries, MongoDB is sufficient and performs better than postgresql for simple queries
- MongoDB's flexible schema in comparison to SQL where the schema is rigid and needs to be defined upfront