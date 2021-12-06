package main

import (
	"fmt"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestMain(t *testing.T) {
	app := cli.NewApp()

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "folder",
			Value: "folder",
			Usage: "Folder name",
			// Destination: &folderName,
		},
	}
	folder := &cli.Command{
		Name:  "new",
		Usage: "folder",
		Flags: myFlags,
		Action: func(c *cli.Context) error {
			//Project
			folderName := c.String("folder")
			fmt.Println(folderName)
			return nil
		},
	}

	app.Commands = cli.Commands{folder}
	err := app.Run([]string{"gojira", "new"})

	if err != nil {
		t.Errorf("should not return an error %v", err)
	}
}
