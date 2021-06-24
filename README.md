# Gojira
## _Create project with already base app configuration_

[![N|Solid](https://golang.org/lib/godoc/images/go-logo-blue.svg)](https://nodesource.com/products/nsolid)


Gojira is a cli tool to create base app for you.

- Creates base config project for you

## Features

- Gin Gonic, Creates a gin gonic app with Jwt, Bcrypt and go mod


## Installation

Gojira requires [Go](https://golang.org/) v1.11+ to run.

Install the dependencies and devDependencies and start the server.

```sh
go get github.com/saturnavt/gojira
```


### How to use

In your terminal type:

```sh
gojira // This will show all the commands available
```

To create a new gin-gonic project:

```
gojira gin --folder yourProjectName // yourProjectName is the name of your new project
```

## License

MIT

Gojira is [MIT licensed](LICENSE).
