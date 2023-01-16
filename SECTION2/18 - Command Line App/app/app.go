package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
	CliPackage "github.com/urfave/cli"
)

// This function will return the Command Line Application.
func Generate() *CliPackage.App {
	app := CliPackage.NewApp()

	app.Name = "Command Line Application"
	app.Usage = "Search for names and servers IP's."

	app.Commands = []CliPackage.Command{
		{
			Name:  "ip",
			Usage: "search for ips in web",
			Flags: []CliPackage.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *CliPackage.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Printf("Public IP: %v\n", ip)
	}
}
