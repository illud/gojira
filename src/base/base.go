package base

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

func BaseData(folderName string) {
	//Add data to main.go
	mainString :=
		`package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	tasksController "github.com/` + folderName + `/controller/tasks"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	//tasks
	router.POST("/tasks", tasksController.CreateTasks)
	router.GET("/tasks", tasksController.GetTasks)
	router.PUT("/tasks/:id", tasksController.UpdateTasks)
	router.DELETE("/tasks/:id", tasksController.DeleteTasks)
	
	router.Run(":5000")
}`
	mainBytes := []byte(mainString)
	ioutil.WriteFile(folderName+"/main.go", mainBytes, 0)

	//Add data to task-controller.go
	taskControllerString :=
		`package tasks

import (
	"github.com/gin-gonic/gin"
	tasksUseCase "github.com/` + folderName + `/domain/useCase/tasks"
	tasksEntity "github.com/gojira/infraestructure/entities/tasks"
)

func CreateTasks(c *gin.Context) {
	var task tasksEntity.Task
	c.ShouldBindJSON(&task)

	c.JSON(200, gin.H{
		"message": tasksUseCase.CreateTasks(task.Id, task.Title, task.Description),
	})
}

func GetTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": tasksUseCase.GetTasks(),
	})
}

func UpdateTasks(c *gin.Context) {
	var task tasksEntity.Task
	c.ShouldBindJSON(&task)
	taskId := c.Param("id")

	c.JSON(200, gin.H{
		"message": tasksUseCase.UpdateTasks(taskId, task.Title),
	})
}

func DeleteTasks(c *gin.Context) {
	taskId := c.Param("id")

	c.JSON(200, gin.H{
		"message": tasksUseCase.DeleteTasks(taskId),
	})
}`
	taskControllerBytes := []byte(taskControllerString)
	ioutil.WriteFile(folderName+"/controller/tasks/task.controller.go", taskControllerBytes, 0)

	//Add data to task.useCase.go
	taskUseCaseString :=
		`package tasks

import (
	tasksRepository "github.com/` + folderName + `/infraestructure/repository/tasks"
)

func CreateTasks(taskId string, title string, description string) string {
	return tasksRepository.CreateTasks(taskId, title, description)
}

func GetTasks() string {
	return tasksRepository.FindAll()
}

func UpdateTasks(taskId string, title string) string {
	return tasksRepository.UpdateTasks(taskId, title)
}

func DeleteTasks(taskId string) string {
	return tasksRepository.Delete(taskId)
}`
	taskUseCaseBytes := []byte(taskUseCaseString)
	ioutil.WriteFile(folderName+"/domain/useCase/tasks/tasks.useCase.go", taskUseCaseBytes, 0)

	//Add data to task.repository.go
	taskRepositoryString :=
		`package tasks

func CreateTasks(taskId string, title string, description string) string {
	return "task created " + taskId + " " + title + " " + description
}

func FindAll() string {
	return "{id: 1, title: 'Task title', description: 'Task description'}"
}

func UpdateTasks(taskId string, title string) string {
	return "task updated " + taskId + " " + title
}

func Delete(taskId string) string {
	return "task deleted " + taskId
}`
	taskRepositoryBytes := []byte(taskRepositoryString)
	ioutil.WriteFile(folderName+"/infraestructure/repository/tasks/tasks.repository.go", taskRepositoryBytes, 0)

	//Add data to task.entity.go
	taskEntitiesString :=
		`package tasks

type Task struct {
	Id          string
	Title       string
	Description string
}`
	taskEntitiesBytes := []byte(taskEntitiesString)
	ioutil.WriteFile(folderName+"/infraestructure/entities/tasks/tasks.entity.go", taskEntitiesBytes, 0)

	//Add data to bcrypt.go
	bcrypt := `package services

import (
	"golang.org/x/crypto/bcrypt"
)

//hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

//check password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}`

	bcryptBytes := []byte(bcrypt)
	ioutil.WriteFile(folderName+"/utils/services/bcrypt/bcrypt.go", bcryptBytes, 0)

	// JWT
	jwtString :=
		`package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString
}

func ValidateToken(validate string) string {
	var tokenCheker string
	tokenString := validate
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	// ... error handling
	if err != nil {
		fmt.Println("Error: ", err)
		tokenCheker = "Error"
	} else {
		tokenCheker = "Ok"
	}

	return tokenCheker
}`
	//Add data to jwt.go
	jwtBytes := []byte(jwtString)
	ioutil.WriteFile(folderName+"/utils/services/jwt/jwt.go", jwtBytes, 0)

	// ERRORS
	errorsString :=
		`package errors`
	//Add data to errors.go
	errorsBytes := []byte(errorsString)
	ioutil.WriteFile(folderName+"/utils/errors/errors.go", errorsBytes, 0)

	// Add database client
	clientString :=
		`package databases`
	//Add data to client.go
	clientBytes := []byte(clientString)
	ioutil.WriteFile(folderName+"/infraestructure/databases/client.go", clientBytes, 0)
}

func BaseModuleCrud(moduleName string) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	//Add data to controller.go
	controllerString :=
		`package ` + moduleName + `

import (
	"github.com/gin-gonic/gin"
	` + moduleName + `UseCase "github.com/` + currentDirName + `/domain/useCase/` + moduleName + `"
	` + moduleName + `Entity "github.com/gojira/infraestructure/entities/` + moduleName + `"
)

func Create` + strings.Title(strings.ToLower(moduleName)) + `(c *gin.Context) {
	var ` + moduleName + ` ` + moduleName + `Entity.` + strings.Title(strings.ToLower(moduleName)) + `
	c.ShouldBindJSON(&` + moduleName + `)

	c.JSON(200, gin.H{
		"message": ` + moduleName + `UseCase.Create` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `.Id, ` + moduleName + `.Title),
	})
}

func Get` + strings.Title(strings.ToLower(moduleName)) + `(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": ` + moduleName + `UseCase.Get` + strings.Title(strings.ToLower(moduleName)) + `(),
	})
}

func Update` + strings.Title(strings.ToLower(moduleName)) + `(c *gin.Context) {
	var ` + moduleName + ` ` + moduleName + `Entity.` + strings.Title(strings.ToLower(moduleName)) + `
	c.ShouldBindJSON(&` + moduleName + `)
	` + moduleName + `Id := c.Param("id")

	c.JSON(200, gin.H{
		"message": ` + moduleName + `UseCase.Update` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id, ` + moduleName + `.Title),
	})
}

func Delete` + strings.Title(strings.ToLower(moduleName)) + `(c *gin.Context) {
	` + moduleName + `Id := c.Param("id")

	c.JSON(200, gin.H{
		"message": ` + moduleName + `UseCase.Delete` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id),
	})
}`
	controllerBytes := []byte(controllerString)
	ioutil.WriteFile("controller/"+moduleName+"/"+moduleName+".controller.go", controllerBytes, 0)

	//Add data to useCase.go
	useCaseString :=
		`package ` + moduleName + `

import (
	` + moduleName + `Repository "github.com/` + currentDirName + `/infraestructure/repository/` + moduleName + `"
)

func Create` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id string, title string) string {
	return ` + moduleName + `Repository.Create` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id, title)
}

func Get` + strings.Title(strings.ToLower(moduleName)) + `() string {
	return ` + moduleName + `Repository.FindAll()
}

func Update` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id string, title string) string {
	return ` + moduleName + `Repository.Update` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id, title)
}

func Delete` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id string) string {
	return ` + moduleName + `Repository.Delete(` + moduleName + `Id)
}`
	useCaseBytes := []byte(useCaseString)
	ioutil.WriteFile("domain/useCase/"+moduleName+"/"+moduleName+".useCase.go", useCaseBytes, 0)

	//Add data to repository.go
	repositoryString :=
		`package ` + moduleName + `

func Create` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id string, title string) string {
	return "` + moduleName + ` created " + ` + moduleName + `Id + " " + title
}

func FindAll() string {
	return "{id: 1, title: '` + moduleName + ` title'}"
}

func Update` + strings.Title(strings.ToLower(moduleName)) + `(` + moduleName + `Id string, title string) string {
	return "` + moduleName + ` updated " + ` + moduleName + `Id + " " + title
}

func Delete(` + moduleName + `Id string) string {
	return "` + moduleName + ` deleted " + ` + moduleName + `Id
}`
	repositoryBytes := []byte(repositoryString)
	ioutil.WriteFile("infraestructure/repository/"+moduleName+"/"+moduleName+".repository.go", repositoryBytes, 0)

	//Add data to moduleName.entity.go
	entitiesString :=
		`package ` + moduleName + `

type ` + strings.Title(strings.ToLower(moduleName)) + ` struct {
	Id    string
	Title string
}`
	entitiesBytes := []byte(entitiesString)
	ioutil.WriteFile("infraestructure/entities/"+moduleName+"/"+moduleName+".entity.go", entitiesBytes, 0)
}

func BaseModuleSimple(moduleName string) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]

	//Add data to controller.go
	controllerString :=
		`package ` + moduleName + `

import (
	"github.com/gin-gonic/gin"
	` + moduleName + `UseCase "github.com/` + currentDirName + `/domain/useCase/` + moduleName + `"
	_ "github.com/gojira/infraestructure/entities/` + moduleName + `" // Change _ for ` + moduleName + `Entity or something that works for you
)

func Get` + strings.Title(strings.ToLower(moduleName)) + `(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": ` + moduleName + `UseCase.Get` + strings.Title(strings.ToLower(moduleName)) + `(),
	})
}`
	controllerBytes := []byte(controllerString)
	ioutil.WriteFile("controller/"+moduleName+"/"+moduleName+".controller.go", controllerBytes, 0)

	//Add data to useCase.go
	useCaseString :=
		`package ` + moduleName + `

import (
	` + moduleName + `Repository "github.com/` + currentDirName + `/infraestructure/repository/` + moduleName + `"
)

func Get` + strings.Title(strings.ToLower(moduleName)) + `() string {
	return ` + moduleName + `Repository.FindAll()
}`
	useCaseBytes := []byte(useCaseString)
	ioutil.WriteFile("domain/useCase/"+moduleName+"/"+moduleName+".useCase.go", useCaseBytes, 0)

	//Add data to repository.go
	repositoryString :=
		`package ` + moduleName + `

func FindAll() string {
	return "{id: 1, title: '` + moduleName + ` title'}"
}`
	repositoryBytes := []byte(repositoryString)
	ioutil.WriteFile("infraestructure/repository/"+moduleName+"/"+moduleName+".repository.go", repositoryBytes, 0)

	//Add data to moduleName.entity.go
	entitiesString :=
		`package ` + moduleName + `

type ` + strings.Title(strings.ToLower(moduleName)) + ` struct {

}`
	entitiesBytes := []byte(entitiesString)
	ioutil.WriteFile("infraestructure/entities/"+moduleName+"/"+moduleName+".entity.go", entitiesBytes, 0)
}

func BaseDbClient(clientName string) {
	// Add database client
	clientString := ""
	if clientName == "mysql" {
		clientString =
			`package databases
		
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var CLIENT = DbConnection()

func DbConnection() *sql.DB {
	//CONNECTION
	db, err := sql.Open("mysql", "databaseUsername:databasePassword@tcp(localhost:3306)/yourDatabaseTablename")
	
	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}
	fmt.Println("CONNECTED")
	return db
}`
	}

	if clientName == "gorm" {
		clientString =
			`package databases
		
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var CLIENT = DbConnection()

func DbConnection() *gorm.DB {
	//CONNECTION
	db, err := gorm.Open("mysql", "databaseUsername:databasePassword@tcp(127.0.0.1:3306)/yourDatabaseTablename?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}
	// sqlDB, err := db.DB()
	// defer sqlDB.Close()
	// defer db.Close()
	fmt.Println("CONNECTED")
	return db
}`
	}

	if clientName == "prisma" {

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
		var ss []string
		if runtime.GOOS == "windows" {
			ss = strings.Split(dir, "\\")
		} else {
			ss = strings.Split(dir, "/")
		}

		currentDirName := ss[len(ss)-1]

		clientString =
			`package databases
		
import (
	"fmt"

	"github.com/` + currentDirName + `/infraestructure/databases/prisma/db"
	"golang.org/x/net/context"
)

var Client = DB()

func DB() *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println(err)
	}

	// defer func() {
	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return client
}

var Context = ContextService()

func ContextService() context.Context {
	ctx := context.Background()
	return ctx
}`

		//Insertdata into prisma.schema
		prismaString :=
			`datasource db {
	// could be postgresql or mysql
	provider = "sqlite"
	url      = "file:dev.db"
}

generator db {
	provider = "go run github.com/prisma/prisma-client-go"
	// set the output folder and package name
	   output           = "./infraestructure/databases/prisma/db"
	   package          = "db"
}

//This is and example table add your own schemas
model Tasks {
	id        Int      @id @default(autoincrement())
	createdAt DateTime @default(now())
	updatedAt DateTime @updatedAt
	title     String
	desc      String
}`

		prismaSChemaBytes := []byte(prismaString)
		ioutil.WriteFile("schema.prisma", prismaSChemaBytes, 0)
	}

	//Add data to client.go
	clientBytes := []byte(clientString)
	ioutil.WriteFile("infraestructure/databases/client.go", clientBytes, 0)
}
