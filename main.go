package main

import "github.com/codegangsta/cli"
import "os"

func main() {
	app := cli.NewApp()
	app.Name = "robot48"
	app.Author = "haoxin"
	app.Version = "0.0.1"
	app.Usage = "Oh, snh48 robot! 😄"

	app.Commands = []cli.Command{
		{
			Name:    "ticket",
			Aliases: []string{"t"},
			Usage:   "query ticket info - robot48 ticket n 2016-02",
			Action: func(c *cli.Context) {
				team, date := "", ""
				args := c.Args()

				if len(args) > 0 {
					team = args[0]
				}
				if len(args) > 1 {
					date = args[1]
				}

				queryTickets(team, date)
			},
		},
	}

	app.Run(os.Args)
}
