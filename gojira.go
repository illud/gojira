package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = `
	 _____        _ _
	/ ____|      (_|_)
	| |  __  ___  _ _ _ __ __ _
	| | |_ |/ _ \| | | '__/ _  |
	| |__| | (_) | | | | | (_| |
	 \_____|\___/| |_|_|  \__,_|
					    _/ |
					    |__/`

	app.Usage = "Alejandro Castillo"
	app.Description = "Gojira is a cli tool to create clean architecture app for you including gin-gonic, bcrypt and jwt."

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "folder",
			Value: "folder",
			Usage: "Folder name",
			// Destination: &folderName,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "new",
			Usage: "To create a new project use the next command (gojira new --folder yourProjectName) this will generate a new gin-gonic project with clean architecture, jwt and bcrypt in services folder ready to use",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				//Project
				folderName := c.String("folder")
				os.MkdirAll(folderName, os.ModePerm)
				fmt.Print(folderName)
				fmt.Println(" Ok")
				os.Create(folderName + "/main.go")
				fmt.Print(folderName + "/main.go")
				fmt.Println(" Ok")

				//Controller
				os.MkdirAll(folderName+"/controller", os.ModePerm)
				fmt.Print(folderName + "/controller")
				fmt.Println(" Ok")

				//Controller/Tasks
				os.MkdirAll(folderName+"/controller/tasks", os.ModePerm)
				fmt.Print(folderName + "/controller/tasks")
				fmt.Println(" Ok")

				//Controller/Tasks/task-controller.go
				os.Create(folderName + "/controller/tasks/task-controller.go")
				fmt.Print(folderName + "/controller/tasks/task-controller.go")
				fmt.Println(" Ok")

				//Domain
				os.MkdirAll(folderName+"/domain", os.ModePerm)
				fmt.Print(folderName + "/domain")
				fmt.Println(" Ok")

				//useCase
				os.MkdirAll(folderName+"/domain/useCase", os.ModePerm)
				fmt.Print(folderName + "/domain/useCase")
				fmt.Println(" Ok")

				//useCase/tasks
				os.MkdirAll(folderName+"/domain/useCase/tasks", os.ModePerm)
				fmt.Print(folderName + "/domain/useCase/tasks")
				fmt.Println(" Ok")

				//useCase/tasks/tasks-useCase.go
				os.Create(folderName + "/domain/useCase/tasks/tasks-useCase.go")
				fmt.Print(folderName + "/domain/useCase/tasks/tasks-useCase.go")
				fmt.Println(" Ok")

				//Infraestructure
				os.MkdirAll(folderName+"/infraestructure", os.ModePerm)
				fmt.Print(folderName + "/infraestructure")
				fmt.Println(" Ok")

				//Databases
				os.MkdirAll(folderName+"/infraestructure/databases", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/databases")
				fmt.Println(" Ok")

				//Databases/client.go
				os.Create(folderName + "/infraestructure/databases/client.go")
				fmt.Print(folderName + "/infraestructure/databases/client.go")
				fmt.Println(" Ok")

				//Repository
				os.MkdirAll(folderName+"/infraestructure/repository", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/repository")
				fmt.Println(" Ok")

				//Repository/tasks
				os.MkdirAll(folderName+"/infraestructure/repository/tasks", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/repository/tasks")
				fmt.Println(" Ok")

				//Repository/tasks/tasks-repository.go
				os.Create(folderName + "/infraestructure/repository/tasks/tasks-repository.go")
				fmt.Print(folderName + "/infraestructure/repository/tasks/tasks-repository.go")
				fmt.Println(" Ok")

				//Utils
				os.MkdirAll(folderName+"/infraestructure/utils", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/utils")
				fmt.Println(" Ok")

				//Utils/Errors
				os.MkdirAll(folderName+"/infraestructure/utils/errors", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/utils/errors")
				fmt.Println(" Ok")

				//utils/errors/errors.go
				os.Create(folderName + "/infraestructure/utils/errors/errors.go")
				fmt.Print(folderName + "/infraestructure/utils/errors/errors.go")
				fmt.Println(" Ok")

				//SERVICES
				os.MkdirAll(folderName+"/infraestructure/utils/services", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/utils/services")
				fmt.Println(" Ok")

				//SERVICES/Jwt
				os.MkdirAll(folderName+"/infraestructure/utils/services/jwt", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/utils/services/jwt")
				fmt.Println(" Ok")

				//SERVICES/Jwt/jwt.go
				os.Create(folderName + "/infraestructure/utils/services/jwt/jwt.go")
				fmt.Print(folderName + "/infraestructure/utils/services/jwt/jwt.go")
				fmt.Println(" Ok")

				//SERVICES/bcrypt
				os.MkdirAll(folderName+"/infraestructure/utils/services/bcrypt", os.ModePerm)
				fmt.Print(folderName + "/infraestructure/utils/services/bcrypt")
				fmt.Println(" Ok")

				//SERVICES/bcrypt/bcrypt.go
				os.Create(folderName + "/infraestructure/utils/services/bcrypt/bcrypt.go")
				fmt.Print(folderName + "/infraestructure/utils/services/bcrypt/bcrypt.go")
				fmt.Println(" Ok")

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

	// // do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)

	// }
	return tokenCheker

}`
					//Add data to jwt.go
				jwtBytes := []byte(jwtString)
				ioutil.WriteFile(folderName+"/infraestructure/utils/services/jwt/jwt.go", jwtBytes, 0)

				fmt.Println("")
				fmt.Println(" | To start ")
				fmt.Println(" | cd ", folderName)
				fmt.Println(" | go run main.go ")

				if runtime.GOOS == "windows" {
					cmd := exec.Command("cmd", "/c", "go mod init github.com/"+folderName)
					cmd.Dir = folderName

					//INSTALL DEPENDENCIES
					out, err := cmd.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(out))

					installDependencies := exec.Command("cmd", "/c", "go get -d ./...")
					installDependencies.Dir = folderName

					//INSTALL DEPENDENCIES
					installDependenciesOut, err := installDependencies.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(installDependenciesOut))
				}

				if runtime.GOOS == "linux" {
					cmd := exec.Command("cmd", "/c", "go mod init github.com/"+folderName)
					cmd.Dir = folderName

					//INSTALL DEPENDENCIES
					out, err := cmd.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(out))

					installDependencies := exec.Command("cmd", "/c", "go get -d ./...")
					installDependencies.Dir = folderName

					//INSTALL DEPENDENCIES
					installDependenciesOut, err := installDependencies.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(installDependenciesOut))
				}

				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
