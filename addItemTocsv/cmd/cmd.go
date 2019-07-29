package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

func do() error {
	if filename == "" {
		return errors.New("File no exist")
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	csvr := csv.NewReader(file)

	newRows := make([][]string, 0)

	for {
		r, err := csvr.Read()
		if err == io.EOF {
			break
		}

		switch action {
		case Insert:
			r = insertRows(r, rowIndex, str)
		case Remove:
			r = removeRows(r, rowIndex)
		default:
			return errors.New("The Wrong Action ")
		}

		newRows = append(newRows, r)

	}

	if output == filename {
		return errors.New("have the same filename")
	}

	of, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer of.Close()

	w := csv.NewWriter(of)

	w.WriteAll(newRows)

	w.Flush()

	fmt.Printf("Total Rows : %v", len(newRows))

	return nil
}

func insertRows(s []string, index int, str string) []string {
	ly := len(s) - 1
	if index > ly {
		s = append(s, make([]string, index-ly)...)
		s[index] = str
		return s
	}

	y := s[index:ly]
	s[index] = str
	s = append(s, y...)

	return s
}

func removeRows(s []string, index int) []string {
	return nil
}
