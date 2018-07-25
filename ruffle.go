package main

import (
	"log"
	"os"
	"time"

	"github.com/SamuelTissot/ruffle/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "Ruffle"
	app.Usage = "a break in the smoothness or evenness of data; undulation of content"
	app.Compiled = time.Now()
	app.Copyright = "MIT License"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Samuel Tissot",
			Email: "tissotjobin@gmail.com",
		},
	}

	// global flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "pretty, p",
			Usage: "prettify the SHA256 by encoding it in base64",
		},
	}

	app.Commands = []cli.Command{
		cmd.Encrypt,
		cmd.Find,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
