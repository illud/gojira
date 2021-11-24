# Gojira
## Create project with Clean Architecture folder structure

[![N|Solid](https://golang.org/lib/godoc/images/go-logo-blue.svg)](https://nodesource.com/products/nsolid)


Gojira is a cli tool to create clean architecture app for you including gin-gonic, bcrypt and jwt.

- Creates Clean Architecture project for you

## Features
- Clean Architecture Folder Structure (https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- Gin Gonic (https://github.com/gin-gonic/gin)
- Jwt (https://github.com/dgrijalva/jwt-go)
- Bcrypt (https://golang.org/x/crypto/bcrypt)
- Example tasks api

## Installation

Gojira requires [Go](https://golang.org/) v1.11+ to run.

Install the dependencies.

```sh
go get github.com/saturnavt/gojira
```


### How to use

In your terminal type:

```sh
gojira // This will show all the commands available
```

To create a new gin-gonic with clean architecture project:

```
gojira new --folder yourProjectName
```

To create a new module:

```
gojira module --generate yourModuleName
```

Folder Structure:

```

|-- controller
|        |      
|        |-tasks --> This is and example package(Create your own packages)
|            |       
|            |-tasks-controller.go --> This is and example controller file(Create your own controllers)
|            
|-- domain
|      |
|      |-useCase
|            |
|            |-tasks --> This is and example package(Create your own packages)
|                | 
|                |-task-useCase.go --> This is and example useCase file(Create your own useCases)
|
|-- infraestructure
            |
            |-databases
            |      |
            |      |-client.go
            |
            |-respository
            |      |
            |      |-tasks  --> This is and example package(Create your own packages)
            |          | 
            |          |-tasks-repository.go --> This is and example repository file(Create your own repositories)
            |-utils
                |
                |-errors
                |    |
                |    |-errros.go
                |
                |-services
                     |
                     |-jwt
                     |  |
                     |  |-jwt.go
                     |
                     |-bcrypt
                        |
                        |-bcrypt.go
```

## License

MIT

Gojira is [MIT licensed](LICENSE).
