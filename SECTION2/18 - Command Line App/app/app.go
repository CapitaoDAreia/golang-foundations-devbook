package app

import CliPackage "github.com/urfave/cli"

//This function will return the Command Line Application.
func Generate() *CliPackage.App {
	app := CliPackage.NewApp()

	app.Name = "Command Line Application"
	app.Usage = "Search for names and servers IP's."

	return app
}
