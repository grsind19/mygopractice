package main

import (
	"go/parser"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Calculator"
	app.Usage = "Common math calculations"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "add",
			Value: "see",
			Usage: "Add two numbers",
		},
	}
	app.Action = func(c *cli.Context) err {
		var params = c.GlobalString("add")
		exp := parser.ParseExpr(params)
		return exp
	}
	app.Run(os.Args)

}
