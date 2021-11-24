package base

import (
	"io/ioutil"
)

func BaseData(folderName string) {
	//Add data to main.go
	mainString :=
		`package main

import (
	"github.com/gin-gonic/gin"
	tasksController "github.com/` + folderName + `/controller/tasks"
)

func main() {
	router := gin.Default()
	//tasks
	router.GET("/tasks", tasksController.GetTasks)
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
)

func GetTasks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": tasksUseCase.GetTasks(),
	})
}`
	taskControllerBytes := []byte(taskControllerString)
	ioutil.WriteFile(folderName+"/controller/tasks/task-controller.go", taskControllerBytes, 0)

	//Add data to task-useCase.go
	taskUseCaseString :=
		`package tasks

import (
	tasksRepository "github.com/` + folderName + `/infraestructure/repository/tasks"
)

func GetTasks() string {
	return tasksRepository.FindAll()
}`
	taskUseCaseBytes := []byte(taskUseCaseString)
	ioutil.WriteFile(folderName+"/domain/useCase/tasks/tasks-useCase.go", taskUseCaseBytes, 0)

	//Add data to task-repository.go
	taskRepositoryString :=
		`package tasks

func FindAll() string {
	return "Hello. Gojira"
}`
	taskRepositoryBytes := []byte(taskRepositoryString)
	ioutil.WriteFile(folderName+"/infraestructure/repository/tasks/tasks-repository.go", taskRepositoryBytes, 0)

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
	ioutil.WriteFile(folderName+"/infraestructure/utils/services/bcrypt/bcrypt.go", bcryptBytes, 0)

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
	ioutil.WriteFile(folderName+"/infraestructure/utils/services/jwt/jwt.go", jwtBytes, 0)

	// ERRORS
	errorsString :=
		`package errors`
	//Add data to errors.go
	errorsBytes := []byte(errorsString)
	ioutil.WriteFile(folderName+"/infraestructure/utils/errors/errors.go", errorsBytes, 0)

	// Add database client
	clientString :=
		`package databases`
	//Add data to client.go
	clientBytes := []byte(clientString)
	ioutil.WriteFile(folderName+"/infraestructure/databases/client.go", clientBytes, 0)
}
