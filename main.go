package main

import "github.com/codegangsta/cli"
import "time"
import "fmt"
import "log"
import "os"

func main() {
	app := cli.NewApp()
	app.Name = "robot48"
	app.Author = "haoxin"
	app.Version = "0.1.0"
	app.Usage = "Oh, snh48 robot! 😄"

	app.Commands = []cli.Command{
		{
			Name:    "ticket",
			Aliases: []string{"T"},
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
		}, {
			Name:    "profile",
			Aliases: []string{"P"},
			Usage:   "query profile - robot48 profile",
			Action: func(c *cli.Context) {
				queryProfile()
			},
		}, {
			Name:    "order",
			Aliases: []string{"O"},
			Usage:   "get order list - robot48 order",
			Action: func(c *cli.Context) {
				getOrders()
			},
		}, {
			Name:    "signin",
			Aliases: []string{"S"},
			Usage:   "signin",
			Action: func(c *cli.Context) {
				wl := newWsLiner()
				defer wl.Close()

				name, e := wl.Prompt("username:")
				if e != nil {
					log.Fatal(e)
				}
				pass, e := wl.PasswordPrompt("password:")
				if e != nil {
					log.Fatal(e)
				}

				fmt.Println(name, pass)
			},
		}, {
			Name:    "repl",
			Aliases: []string{"R"},
			Usage:   "repl mode",
			Action: func(c *cli.Context) {
				wl := newWsLiner()
				defer wl.Close()

				for {
					time.Sleep(100 * time.Microsecond)

					i := wl.readInput()
					if i != "" {
						fmt.Println("cmd> " + i)
					}
				}
			},
		},
	}

	app.Run(os.Args)
}
