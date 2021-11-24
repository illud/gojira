package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	base "github.com/saturnavt/gojira/src/base"
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
		&cli.StringFlag{
			Name:  "generate",
			Value: "generate",
			Usage: "Generates module",
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

				//Create base files data
				base.BaseData(folderName)

				//Display usage
				fmt.Println("")
				fmt.Println(" | get started ")
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
		{
			Name:  "module",
			Usage: "To create a new module use the next command (gojira module --generate yourModuleName) this will generate a new module",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				// dir, err := os.Getwd()
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// fmt.Println(dir)
				// var ss []string
				// if runtime.GOOS == "windows" {
				// 	ss = strings.Split(dir, "\\")
				// } else {
				// 	ss = strings.Split(dir, "/")
				// }

				// currentDirName := ss[len(ss)-1]

				//module name
				moduleName := c.String("generate")

				//Controller/currentDirName
				os.MkdirAll("controller/"+moduleName, os.ModePerm)
				fmt.Print("controller/" + moduleName)
				fmt.Println(" Ok")

				//Controller/moduleName/moduleName-controller.go
				os.Create("controller/" + moduleName + "/" + moduleName + "-controller.go")
				fmt.Print("controller/" + moduleName + "/" + moduleName + "-controller.go")
				fmt.Println(" Ok")

				//useCase
				//useCase/moduleName
				os.MkdirAll("domain/useCase/"+moduleName, os.ModePerm)
				fmt.Print("domain/useCase/" + moduleName)
				fmt.Println(" Ok")

				//useCase/moduleName/moduleName-useCase.go
				os.Create("domain/useCase/" + moduleName + "/" + moduleName + "-useCase.go")
				fmt.Print("domain/useCase/" + moduleName + "/" + moduleName + "-useCase.go")
				fmt.Println(" Ok")

				//Repository
				//Repository/moduleName
				os.MkdirAll("infraestructure/repository/"+moduleName, os.ModePerm)
				fmt.Print("infraestructure/repository/" + moduleName)
				fmt.Println(" Ok")

				//Repository/moduleName/moduleName-repository.go
				os.Create("infraestructure/repository/" + moduleName + "/" + moduleName + "-repository.go")
				fmt.Print("infraestructure/repository/" + moduleName + "/" + moduleName + "-repository.go")
				fmt.Println(" Ok")

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
