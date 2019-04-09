package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"

	"github.com/micheam/gen-img/command"
)

const version = "0.1.0"
const appName = "gen-img"

func main() {
	err := newApp().Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "Generate image with spesified format, size."
	app.Version = version
	app.Author = "Michto Maeda"
	app.Email = "michito.maeda@gmail.com"

	app.Action = command.DoGenerate
	app.ArgsUsage = "[out-file-name]"
	app.Flags = command.RootFlag

	app.Description = `
    Generate image with specified format.
    If it can be determined from the out-file extension,
    specification of the format can be omitted.`

	return app
}
