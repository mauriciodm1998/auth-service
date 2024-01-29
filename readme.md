# FIAP - TechChallenge - Auth Service

## Description

This service is responsable to generate tokens to have authorization to send requests to other services. In this proccess, this service takes the login inputed and searches the user in the postgres-user database, after validation, it saves the access in a table of access and returns the token or an error. We have a diagram about a flow of this service here: [login flow](./docs/diagrams/auth-service-diagram.png), [bypass flow](./docs/diagrams/bypass-diagram.png)

## Features

- Generate Token
- Generate Guest Token

## How To Run Locally

First of all we need the DataBase. The database for this application is shared with user-service; the steps to run this database are described [there](https://github.com/mauriciodm1998/user-service/blob/master/readme.md).
Then you can run the application:

### VSCode - Debug
The launch.json file is already configured for debuging. Just hit F5 and be happy.

### Running directly from go

Option 1: $```go run cmd/auth/main.go```

Option 2: $```make run-app```

## Manually testing the API

On directory ```/api``` there's a collection that can be imported on Insomnia or similar so you can test manually the application's API.

## Running the unit tests

Simply run ```make run-tests``` and let the magic happens. At the end it will automatically open an html with the coverage % for every package.
We also have the most recently applied unit tests file in this [folder](./docs/unit-tests-results/unit-tests.png) too.

## Test + Build + Bake Image

Simply run ```make test-build-bake``` and let the magic happens. The docker file will run the unit-tests, build the application and bake the docker image for the application.

## Infrastructure

This application runs in as a lambda. The terraform about the configuration of this application are in this [repository](https://github.com/mauriciodm1998/tech-challenge-gitops).