package cmd

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func process(path string, skip bool, f func(record []string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)

	if skip {
		// disregard first row
		_, err := r.Read()
		if err == io.EOF {
			return err
		}
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = f(record)
		if err != nil {
			return err
		}
	}
	return nil
}

func hash(v []byte) string {
	h := sha256.Sum256(v)
	return base64.StdEncoding.EncodeToString(h[:])
}

func checkIndexOutOfBound(i int, r []string) error {
	if i > len(r)-1 {
		return fmt.Errorf("\n\ncell index out of bound. num of cell: %d\nNOTE: the index is zero based, 0, 1, 2, ...\n\n", len(r)-1)
	}
	return nil
}
