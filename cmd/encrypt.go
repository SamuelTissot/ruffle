package cmd

import (
	"encoding/csv"
	"os"

	"github.com/urfave/cli"
)

var Encrypt = cli.Command{
	Name:  "encrypt",
	Usage: "perform an encryption operation on each line of a `FILE`",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "cell, c",
			Usage: "the index of the cell to operate on, (0 based)",
			Value: 0,
		},
		cli.BoolFlag{
			Name:  "headerRow, hr",
			Usage: "if the first row a header file",
		},
	},
	Action: encrypt,
}

func encrypt(c *cli.Context) error {
	path := c.Args().First()
	w := csv.NewWriter(os.Stdout)
	i := c.Int("cell")
	skip := c.IsSet("headerRow")
	err := process(path, skip, func(record []string) error {
		err := checkIndexOutOfBound(i, record)
		if err != nil {
			return err
		}

		record[i] = hash([]byte(record[i]))
		if err := w.Write(record); err != nil {
			return err
		}
		return nil
	})
	w.Flush()
	return err
}
