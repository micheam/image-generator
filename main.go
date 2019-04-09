package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"

	"github.com/micheam/image-generator/command"
)

const version = "0.1.0"

func main() {
	err := newApp().Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "image-generator"
	app.Usage = "Generate image with spesified format, size, color."
	app.Version = version
	app.Author = "Michto Maeda"
	app.Email = "michito.maeda@gmail.com"
	app.Commands = commands
	return app
}

var commands = []cli.Command{
	command.CmdGenerate,
}
