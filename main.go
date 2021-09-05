package main

import (
	"os"

	"homework-rakamin-go-sql/app"
	"homework-rakamin-go-sql/cli"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
