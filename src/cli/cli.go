package cli

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/saturnavt/gojira/src/base"
	input "github.com/saturnavt/gojira/src/cli/input"
)

var choices = []string{"New project", "Module", "Module with crud", "DB service"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Choose a option\n\n")
	s.WriteString("up/down: to select\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("[x] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func Command() {

	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
		// fmt.Printf("\n---\nYou chose %s!\n", m.choice)

		if m.choice == "New project" {
			fmt.Printf("\n")
			fmt.Println("Enter Project Name: ")
			folderName := input.Input()
			fmt.Println(folderName)
			//Project

			os.MkdirAll(folderName, os.ModePerm)
			fmt.Print(folderName)
			fmt.Println(" Ok")
			os.Create(folderName + "/main.go")
			fmt.Print(folderName + "/main.go")
			fmt.Println(" Ok")

			//Routing
			os.MkdirAll(folderName+"/routing", os.ModePerm)
			fmt.Print(folderName + "/routing")
			fmt.Println(" Ok")
			os.Create(folderName + "/routing/routing.go")
			fmt.Print(folderName + "/routing/routing.go")
			fmt.Println(" Ok")

			//Controller
			os.MkdirAll(folderName+"/controller", os.ModePerm)
			fmt.Print(folderName + "/controller")
			fmt.Println(" Ok")

			//Controller/Tasks
			os.MkdirAll(folderName+"/controller/tasks", os.ModePerm)
			fmt.Print(folderName + "/controller/tasks")
			fmt.Println(" Ok")

			//Controller/Tasks/tasks.controller.go
			os.Create(folderName + "/controller/tasks/tasks.controller.go")
			fmt.Print(folderName + "/controller/tasks/tasks.controller.go")
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
			os.Create(folderName + "/domain/useCase/tasks/tasks.useCase.go")
			fmt.Print(folderName + "/domain/useCase/tasks/tasks.useCase.go")
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

			//Entities
			os.MkdirAll(folderName+"/infraestructure/entities", os.ModePerm)
			fmt.Print(folderName + "/infraestructure/entities")
			fmt.Println(" Ok")

			//Entities/tasks
			os.MkdirAll(folderName+"/infraestructure/entities/tasks", os.ModePerm)
			fmt.Print(folderName + "/infraestructure/entities/tasks")
			fmt.Println(" Ok")

			//Entities/tasks.entity.go
			os.Create(folderName + "/infraestructure/entities/tasks/tasks.entity.go")
			fmt.Print(folderName + "/infraestructure/entities/tasks/tasks.entity.go")
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
			os.Create(folderName + "/infraestructure/repository/tasks/tasks.repository.go")
			fmt.Print(folderName + "/infraestructure/repository/tasks/tasks.repository.go")
			fmt.Println(" Ok")

			//Utils
			os.MkdirAll(folderName+"/utils", os.ModePerm)
			fmt.Print(folderName + "/utils")
			fmt.Println(" Ok")

			//Utils/async
			os.MkdirAll(folderName+"/utils/async", os.ModePerm)
			fmt.Print(folderName + "/utils/async")
			fmt.Println(" Ok")

			//utils/async/async.go
			os.Create(folderName + "/utils/async/async.go")
			fmt.Print(folderName + "/utils/async/async.go")
			fmt.Println(" Ok")

			//Utils/Errors
			os.MkdirAll(folderName+"/utils/errors", os.ModePerm)
			fmt.Print(folderName + "/utils/errors")
			fmt.Println(" Ok")

			//utils/errors/errors.go
			os.Create(folderName + "/utils/errors/errors.go")
			fmt.Print(folderName + "/utils/errors/errors.go")
			fmt.Println(" Ok")

			//SERVICES
			os.MkdirAll(folderName+"/utils/services", os.ModePerm)
			fmt.Print(folderName + "/utils/services")
			fmt.Println(" Ok")

			//SERVICES/Jwt
			os.MkdirAll(folderName+"/utils/services/jwt", os.ModePerm)
			fmt.Print(folderName + "/utils/services/jwt")
			fmt.Println(" Ok")

			//SERVICES/Jwt/jwt.go
			os.Create(folderName + "/utils/services/jwt/jwt.go")
			fmt.Print(folderName + "/utils/services/jwt/jwt.go")
			fmt.Println(" Ok")

			//SERVICES/bcrypt
			os.MkdirAll(folderName+"/utils/services/bcrypt", os.ModePerm)
			fmt.Print(folderName + "/utils/services/bcrypt")
			fmt.Println(" Ok")

			//SERVICES/bcrypt/bcrypt.go
			os.Create(folderName + "/utils/services/bcrypt/bcrypt.go")
			fmt.Print(folderName + "/utils/services/bcrypt/bcrypt.go")
			fmt.Println(" Ok")

			//TEST FOLDER
			os.MkdirAll(folderName+"/test", os.ModePerm)
			fmt.Print(folderName + "/test")
			fmt.Println(" Ok")

			//TEST TASKS FOLDER
			os.MkdirAll(folderName+"/test/tasks", os.ModePerm)
			fmt.Print(folderName + "/test/tasks")
			fmt.Println(" Ok")

			os.Create(folderName + "/test/tasks/getTasks_test.go")
			fmt.Print(folderName + "/test/tasks/getTasks_test.go")
			fmt.Println(" Ok")

			//Create base files data
			base.BaseData(folderName)

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

				//SWAG INIT
				swagInit := exec.Command("cmd", "/c", "swag init")
				swagInit.Dir = folderName

				swagInitOut, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(swagInitOut))

				//INSTALL DEPENDENCIES
				installDependenciesOut, err := installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(installDependenciesOut))

				//INSTALL TEST DEPENDENCIES
				installTestDependencies := exec.Command("cmd", "/c", "go get -t ./...")
				installTestDependencies.Dir = folderName

				installTestDependenciesOut, err := installTestDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(installTestDependenciesOut))

			}

			if runtime.GOOS == "linux" {
				cmd := exec.Command("sh", "/c", "go mod init github.com/"+folderName)
				cmd.Dir = folderName

				//INSTALL DEPENDENCIES
				out, err := cmd.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(out))

				installDependencies := exec.Command("sh", "/c", "go get -d ./...")
				installDependencies.Dir = folderName

				//SWAG INIT
				swagInit := exec.Command("sh", "/c", "swag init")
				swagInit.Dir = folderName

				swagInitOut, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(swagInitOut))

				//INSTALL DEPENDENCIES
				installDependenciesOut, err := installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(installDependenciesOut))

				//INSTALL TEST DEPENDENCIES
				installTestDependencies := exec.Command("sh", "/c", "go get -t ./...")
				installTestDependencies.Dir = folderName

				installTestDependenciesOut, err := installTestDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(installTestDependenciesOut))
			}

			//Display usage
			fmt.Println(" | get started ")
			fmt.Println(" | cd ", folderName)
			fmt.Println(" | go run main.go ")
			fmt.Println("")

		}

		if m.choice == "Module" {
			fmt.Printf("\n")
			fmt.Println("Enter Module Name: ")
			moduleName := input.Input()

			//Controller/currentDirName
			os.MkdirAll("controller/"+moduleName, os.ModePerm)
			fmt.Print("controller/" + moduleName)
			fmt.Println(" Ok")

			//Controller/moduleName/moduleName-controller.go
			os.Create("controller/" + moduleName + "/" + moduleName + ".controller.go")
			fmt.Print("controller/" + moduleName + "/" + moduleName + ".controller.go")
			fmt.Println(" Ok")

			//useCase
			//useCase/moduleName
			os.MkdirAll("domain/useCase/"+moduleName, os.ModePerm)
			fmt.Print("domain/useCase/" + moduleName)
			fmt.Println(" Ok")

			//useCase/moduleName/moduleName-useCase.go
			os.Create("domain/useCase/" + moduleName + "/" + moduleName + ".useCase.go")
			fmt.Print("domain/useCase/" + moduleName + "/" + moduleName + ".useCase.go")
			fmt.Println(" Ok")

			//Repository
			//Repository/moduleName
			os.MkdirAll("infraestructure/repository/"+moduleName, os.ModePerm)
			fmt.Print("infraestructure/repository/" + moduleName)
			fmt.Println(" Ok")

			//Repository/moduleName/moduleName-repository.go
			os.Create("infraestructure/repository/" + moduleName + "/" + moduleName + ".repository.go")
			fmt.Print("infraestructure/repository/" + moduleName + "/" + moduleName + ".repository.go")
			fmt.Println(" Ok")

			//Entities/moduleName
			os.MkdirAll("infraestructure/entities/"+moduleName+"", os.ModePerm)
			fmt.Print("infraestructure/entities/" + moduleName + "")
			fmt.Println(" Ok")

			//Entities/moduleName.entity.go
			os.Create("infraestructure/entities/" + moduleName + "/" + moduleName + ".entity.go")
			fmt.Print("infraestructure/entities/" + moduleName + "/" + moduleName + ".entity.go")
			fmt.Println(" Ok")

			//TEST FOLDER
			os.MkdirAll("test/"+moduleName, os.ModePerm)
			fmt.Print("test/" + moduleName)
			fmt.Println(" Ok")

			os.Create("test/" + moduleName + "/get" + strings.Title(moduleName) + "_test.go")
			fmt.Print("test/" + moduleName + "/get" + strings.Title(moduleName) + "_test.go")
			fmt.Println(" Ok")

			//Generates module with crud data
			base.BaseModuleCrud(moduleName)

			//Append controller to routing.go file
			base.AppendToRoutingCrud(moduleName)

			//SWAG INIT Windows
			if runtime.GOOS == "windows" {
				swagInit := exec.Command("cmd", "/c", "swag init")

				swagInitOut, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(swagInitOut))
			}

			//SWAG INIT Linux
			if runtime.GOOS == "linux" {
				swagInit := exec.Command("sh", "/c", "swag init")

				swagInitOut, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(swagInitOut))
			}
		}

		if m.choice == "Module with crud" {
			fmt.Printf("\n")
			fmt.Println("Enter Module Name: ")
			moduleName := input.Input()

			//Controller/currentDirName
			os.MkdirAll("controller/"+moduleName, os.ModePerm)
			fmt.Print("controller/" + moduleName)
			fmt.Println(" Ok")

			//Controller/moduleName/moduleName-controller.go
			os.Create("controller/" + moduleName + "/" + moduleName + ".controller.go")
			fmt.Print("controller/" + moduleName + "/" + moduleName + ".controller.go")
			fmt.Println(" Ok")

			//useCase
			//useCase/moduleName
			os.MkdirAll("domain/useCase/"+moduleName, os.ModePerm)
			fmt.Print("domain/useCase/" + moduleName)
			fmt.Println(" Ok")

			//useCase/moduleName/moduleName-useCase.go
			os.Create("domain/useCase/" + moduleName + "/" + moduleName + ".useCase.go")
			fmt.Print("domain/useCase/" + moduleName + "/" + moduleName + ".useCase.go")
			fmt.Println(" Ok")

			//Repository
			//Repository/moduleName
			os.MkdirAll("infraestructure/repository/"+moduleName, os.ModePerm)
			fmt.Print("infraestructure/repository/" + moduleName)
			fmt.Println(" Ok")

			//Repository/moduleName/moduleName-repository.go
			os.Create("infraestructure/repository/" + moduleName + "/" + moduleName + ".repository.go")
			fmt.Print("infraestructure/repository/" + moduleName + "/" + moduleName + ".repository.go")
			fmt.Println(" Ok")

			//Entities/moduleName
			os.MkdirAll("infraestructure/entities/"+moduleName+"", os.ModePerm)
			fmt.Print("infraestructure/entities/" + moduleName + "")
			fmt.Println(" Ok")

			//Entities/moduleName.entity.go
			os.Create("infraestructure/entities/" + moduleName + "/" + moduleName + ".entity.go")
			fmt.Print("infraestructure/entities/" + moduleName + "/" + moduleName + ".entity.go")
			fmt.Println(" Ok")

			//TEST FOLDER
			os.MkdirAll("test/"+moduleName, os.ModePerm)
			fmt.Print("test/" + moduleName)
			fmt.Println(" Ok")

			os.Create("test/" + moduleName + "/get" + strings.Title(moduleName) + "_test.go")
			fmt.Print("test/" + moduleName + "/get" + strings.Title(moduleName) + "_test.go")
			fmt.Println(" Ok")

			//Generates module data in file
			base.BaseModuleSimple(moduleName)

			//Append controller to routing.go file
			base.AppendToRoutingSimple(moduleName)

			//SWAG INIT Windows
			if runtime.GOOS == "windows" {
				swagInit := exec.Command("cmd", "/c", "swag init")

				swagInitOut, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(swagInitOut))
			}

			//SWAG INIT Linux
			if runtime.GOOS == "linux" {
				swagInit := exec.Command("sh", "/c", "swag init")

				swagInitOut, err := swagInit.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(swagInitOut))
			}
		}

		if m.choice == "DB service" {
			fmt.Printf("\n")
			fmt.Println("Enter DB(mysql, gorm or prisma) Name: ")
			flagName := input.Input()

			if flagName == "mysql" {
				base.BaseDbClient("mysql")
			}
			if flagName == "gorm" {
				base.BaseDbClient("gorm")
			}
			if flagName == "prisma" {
				base.BaseDbClient("prisma")
				//create/schema.prisma
				os.MkdirAll("infraestructure/databases/prisma/db", os.ModePerm)
				fmt.Print("infraestructure/databases/prisma/db")
				fmt.Println(" Ok")

				os.Create("schema.prisma")
				fmt.Print("schema.prisma")
				fmt.Println(" Ok")

				fmt.Println("")

				fmt.Println("To get this up and running in your database, we use the Prisma migration tool migrate to create and migrate our database:")
				fmt.Println("sync the database with your schema go run github.com/prisma/prisma-client-go migrate dev --name init")
				fmt.Println("After the migration, the Prisma Client Go client is automatically generated in your project. If you just want to re-generate the client, run go run github.com/prisma/prisma-client-go generate.")

				fmt.Println("For more visit https://github.com/prisma/prisma-client-go")

				//Install db DEPENDENCIES
				if runtime.GOOS == "windows" {
					fmt.Println("")
					fmt.Println("executing go get github.com/prisma/prisma-client-go")

					installDependencies := exec.Command("cmd", "/c", "go get github.com/prisma/prisma-client-go")

					//INSTALL DEPENDENCIES
					installDependenciesOut, err := installDependencies.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(installDependenciesOut))

					//Run prisma init
					installPrismaDependencies := exec.Command("cmd", "/c", "go run github.com/prisma/prisma-client-go migrate dev --name init")

					//INSTALL DEPENDENCIES
					installPrismaDependenciesOut, err := installPrismaDependencies.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(installPrismaDependenciesOut))
				}

				if runtime.GOOS == "linux" {
					fmt.Println("")
					fmt.Println("executing go get github.com/prisma/prisma-client-go")

					installDependencies := exec.Command("sh", "/c", "go get github.com/prisma/prisma-client-go.")

					//INSTALL DEPENDENCIES
					installDependenciesOut, err := installDependencies.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(installDependenciesOut))

					//Run prisma init
					installPrismaDependencies := exec.Command("sh", "/c", "go run github.com/prisma/prisma-client-go migrate dev --name init")

					//INSTALL DEPENDENCIES
					installPrismaDependenciesOut, err := installPrismaDependencies.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(installPrismaDependenciesOut))
				}

			}

			//Install db DEPENDENCIES
			if runtime.GOOS == "windows" {
				installDependencies := exec.Command("cmd", "/c", "go get -d ./...")

				//INSTALL DEPENDENCIES
				installDependenciesOut, err := installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(installDependenciesOut))
			}

			if runtime.GOOS == "linux" {
				installDependencies := exec.Command("sh", "/c", "go get -d ./...")

				//INSTALL DEPENDENCIES
				installDependenciesOut, err := installDependencies.Output()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Println(string(installDependenciesOut))
			}
		}
	}
}
