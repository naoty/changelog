package main

import (
	"os"

	"github.com/urfave/cli"
)

// Version is the version of this application.
var Version = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "changelog"
	app.Version = Version
	app.Usage = "Manage CHANGELOG in my style"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Run(os.Args)
}
