package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Counter"
	app.Usage = "Counts up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			Usage:     "Count up",
			ShortName: "u",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Value: 10,
					Usage: "Value count up 10",
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start <= 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := 1; i <= start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			Usage:     "Count dnow",
			ShortName: "d",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Value: 10,
					Usage: "Start counting down from",
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start < 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := 1; i <= start; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
