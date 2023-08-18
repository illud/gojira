package main

import (
	"fmt"

	cmd "github.com/illud/gojira/src/cli"
)

func main() {

	gojira := `
  _____       _ _           
 / ____|     (_|_)          
| |  __  ___  _ _ _ __ __ _ 
| | |_ |/ _ \| | |  __/ _  |
| |__| | (_) | | | | | (_| |
 \_____|\___/| |_|_|  \__ _|
	    _/ |            
	   |__/ Created by Illud
`

	fmt.Println(gojira)
	cmd.Command()
}
