# Gojira

![](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/400px-Go_Logo_Blue.svg.png)
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
- Example tasks api

## Installation

Gojira requires [Go](https://golang.org/) v1.11+ to run.

Install the dependencies.

```sh
go get github.com/saturnavt/gojira
```


### How to use

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
gojira module --generate yourModuleName
```

To create a new module with simple example:

```
gojira module-simple --generate yourModuleName
```

# When you create a module dont forget to add the controller to the main.go file, i'm working so it can be added automatically.



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
            |
|           |-entities
|           |      |
|           |      |-task --> This is and example package(Create your own packages)
            |          |
            |          |-task.entity.go --> This is and example entity file(Create your own entity)
|           |
|           |-respository
|                  |
|                  |-tasks  --> This is and example package(Create your own packages)
|                      | 
|                      |-tasks.repository.go --> This is and example repository file(Create your own repositories)
|                       
|-- utils
        |
        |-errors
        |    |
        |    |-errros.go
        |
        |-services
                |
                |-jwt
                |   |
                |   |-jwt.go
                |
                |-bcrypt
                    |
                    |-bcrypt.go
```

## License

MIT

Gojira is [MIT licensed](LICENSE).
