package cmd

import (
	"encoding/csv"
	"os"

	"github.com/urfave/cli"
)

var Find = cli.Command{
	Name:  "find",
	Usage: "return the row of the record in the `FILE`",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "needle, n",
			Usage: "the needle to find. the `CONTENT`",
		},
		cli.IntFlag{
			Name:  "cell, c",
			Usage: "the index of the cell to operate on, (0 based)",
			Value: 0,
		},
	},
	Action: find,
}

func find(c *cli.Context) error {
	n := c.String("needle")
	i := c.Int("cell")
	hn := hash([]byte(n), c.GlobalBool("pretty"))
	w := csv.NewWriter(os.Stdout)
	path := c.Args().First()
	err := process(path, false, func(record []string) error {
		err := checkIndexOutOfBound(i, record)
		if err != nil {
			return err
		}

		if record[i] == hn {
			if err := w.Write(record); err != nil {
				return err
			}
			return nil
		}
		return nil
	})
	w.Flush()
	return err
}
