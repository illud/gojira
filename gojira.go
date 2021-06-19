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

	app.Usage = "Saturnavt"
	app.Description = "Saturnavt"

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
			Name:  "gin",
			Usage: "use the flag (--folder yourFolderName) to create a new gin-gonic project with jwt and bcrypt in services folder ready to use",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				folderName := c.String("folder")
				os.MkdirAll(folderName, os.ModePerm)
				fmt.Print(folderName)
				fmt.Println(" Ok")
				os.Create(folderName + "/main.go")
				fmt.Print(folderName + "/main.go")
				fmt.Println(" Ok")

				//SERVICES
				os.MkdirAll(folderName+"/services", os.ModePerm)
				fmt.Print(folderName + "/services")
				fmt.Println(" Ok")
				os.Create(folderName + "/services/jwt.go")
				fmt.Print(folderName + "/services/jwt.go")
				fmt.Println(" Ok")
				os.Create(folderName + "/services/bcrypt.go")
				fmt.Print(folderName + "/services/bcrypt.go")
				fmt.Println(" Ok")

				mainString :=
					`package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}`
				mainBytes := []byte(mainString)
				ioutil.WriteFile(folderName+"/main.go", mainBytes, 0)

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
				ioutil.WriteFile(folderName+"/services/bcrypt.go", bcryptBytes, 0)

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
				jwtBytes := []byte(jwtString)
				ioutil.WriteFile(folderName+"/services/jwt.go", jwtBytes, 0)

				if runtime.GOOS == "windows" {
					cmd := exec.Command("cmd", "/c", "go mod init github.com/"+folderName)
					cmd.Dir = folderName

					//INSTALL DEPENDENCIES
					out, err := cmd.Output()
					if err != nil {
						os.Stderr.WriteString(err.Error())
					}
					fmt.Println(string(out))

					installDependencies := exec.Command("cmd", "/c", "go get ./...")
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

					installDependencies := exec.Command("cmd", "/c", "go get ./...")
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
