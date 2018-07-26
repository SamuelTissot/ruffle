package cmd

import (
	"encoding/csv"
	"os"

	"github.com/urfave/cli"
)

var Encrypt = cli.Command{
	Name:  "hash",
	Usage: "perform a hash operation on each line of a `FILE`",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "cell, c",
			Usage: "the index of the cell to operate on, (0 based)",
			Value: 0,
		},
		cli.BoolFlag{
			Name:  "headerRow, -r",
			Usage: "if the first row a header file",
		},
	},
	Action: hashAction,
}

func hashAction(c *cli.Context) error {
	path := c.Args().First()
	w := csv.NewWriter(os.Stdout)
	i := c.Int("cell")
	headerRow := c.IsSet("headerRow")
	err := process(path, func(record []string) error {
		if headerRow {
			headerRow = false
		} else {
			err := checkIndexOutOfBound(i, record)
			if err != nil {
				return err
			}
			record[i] = hash([]byte(record[i]), c.GlobalBool("pretty"))
		}

		if err := w.Write(record); err != nil {
			return err
		}
		return nil
	})
	w.Flush()
	return err
}
