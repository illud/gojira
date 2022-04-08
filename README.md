# Gojira

![](https://raw.githubusercontent.com/saturnavt/eskolvar.github.io/main/assets/img/gojira.png)

[![Test Status](https://github.com/saturnavt/gojira/actions/workflows/go.yml/badge.svg)](https://github.com/saturnavt/gojira/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://pkg.go.dev/badge/github.com/saturnavt/gojira?status.svg)](https://pkg.go.dev/github.com/saturnavt/gojira?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/saturnavt/gojira)](https://goreportcard.com/report/github.com/saturnavt/gojira)
## Create project with Clean Architecture folder structure

\
Gojira is a cli tool to create clean architecture app for you including gin-gonic, bcrypt and jwt.

- Creates Clean Architecture project for you


## Features
- Clean Architecture Folder Structure (https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- Gin Gonic (https://github.com/gin-gonic/gin)
- Jwt (https://github.com/dgrijalva/jwt-go)
- Bcrypt (https://golang.org/x/crypto/bcrypt)
- Auto generate module with crud example
- Auto generate db service client 
  - Mysql
  - Gorm
  - Prisma
- Example tasks api
- Testing (Auto generate test example when creating a new modules)

## Installation

Gojira requires [Go](https://golang.org/) v1.11+ to run.

Install the dependencies.

```sh
go get github.com/saturnavt/gojira
```
Or

```sh
go install github.com/saturnavt/gojira@latest
```
## How to use

In your terminal type to see all avaible commands:

```sh
gojira
```

To create a new gin-gonic with clean architecture project(This includes a crud example with the name of Tasks):

```
gojira new --folder yourProjectName
```

To create a new module with crud:

```
gojira module --generate-crud yourModuleName
```

To create a new module with simple example:

```
gojira module-simple --generate yourModuleName
```

## Database service
To create a new db service client with Mysql, Gorm or Prisma:

Mysql - to learn more visit (https://github.com/go-sql-driver/mysql)
```
gojira db --client mysql
```

Gorm - to learn more visit (https://github.com/jinzhu/gorm)
```
gojira db --client gorm
```

Prisma - to learn more visit (https://github.com/prisma/prisma-client-go)
```
gojira db --client prisma
```

## This will generate a database connection in infraestructure/databases/client.go

<br/>

### For Mysql and Gorm import the service in your repository like for example:

```go
db "github.com/yourProjectName/infraestructure/databases"
```

Example for Mysql:
```go
// Insert new tasks
res, err := db.Client.Exec("INSERT INTO tasks VALUES(DEFAULT, 'Title', 'Desc')")
if err != nil {
  fmt.Println("ERROR: ", err)
}
fmt.Println(res)
```
To learn more visit (https://github.com/go-sql-driver/mysql)

<br/>

Example for Gorm:
```go
// Insert new tasks
err := db.Client.Save(&tasksEntity.Task{
  Title:       "TEST",
  Description: "This is a description",
})

if err != nil {
  fmt.Println(err)
}
```
To learn more visit (https://github.com/jinzhu/gorm)

<br/>

For prisma import:
```go
dbClient "github.com/yourProjectName/infraestructure/databases"
"github.com/yourProjectName/infraestructure/databases/prisma/db"
```
Example for prisma:
```go
// Insert new tasks
createdTask, err := dbClient.Client.Tasks.CreateOne(
  db.Tasks.Title.Set("Hi from Prisma!"),
  db.Tasks.Description.Set("Prisma is a database toolkit and makes databases easy."),
).Exec(dbClient.Context)

if err != nil {
  fmt.Println(err)
}

result, _ := json.MarshalIndent(createdTask, "", "  ")
fmt.Printf("created task: %s\n", result)
```
To learn more visit (https://github.com/prisma/prisma-client-go)

<br/>

Testing: To run tests use
```shell
go test -v ./...
```

Folder Structure:

```

|-- controller
|        |      
|        |-tasks --> This is and example package(Create your own packages)
|            |       
|            |-tasks.controller.go --> This is and example controller file(Create your own controllers)
|            
|-- domain
|      |
|      |-useCase
|            |
|            |-tasks --> This is and example package(Create your own packages)
|                | 
|                |-task.useCase.go --> This is and example useCase file(Create your own useCases)
|
|-- infraestructure
|           |
|           |-databases
|           |      |
|           |      |-client.go
|           |
|           |-entities
|           |      |
|           |      |-tasks --> This is and example package(Create your own packages)
|           |          |
|           |          |-tasks.entity.go --> This is and example entity file(Create your own entity)
|           |
|           |-respository
|                  |
|                  |-tasks --> This is and example package(Create your own packages)
|                      | 
|                      |-tasks.repository.go --> This is and example repository file(Create your own repositories)
|                       
|-- utils
|       |
|       |-errors
|       |    |
|       |    |-errros.go
|       |
|       |-services
|               |
|               |-jwt
|               |   |
|               |   |-jwt.go
|               |
|               |-bcrypt
|                   |
|                   |-bcrypt.go
|
|-- routing
|       |
|       |-routing.go --> This is where al your routes lives
|
|
|-- test --> This is the test folder
|       |
|       |-tasks --> This is a test example folder
|       |   |
|       |   |-getTasks_test.go --> This is a test example file
|       |
|       |-other test folder --> your other test folder
```

## License

MIT

Gojira is [MIT licensed](LICENSE).
