package fileutil

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

// File - This is just placeholder stuff
type File struct {
	FileName string
	Headers  []string
	Rows     [][]string
}

func readCsv(csvfile io.Reader, filename string) *File {

	d2file := &File{FileName: filename}

	reader := csv.NewReader(csvfile)
	reader.Comma = '\t'

	raw, err := reader.ReadAll()
	check(err)

	rows := make([][]string, 0)
	headers := make([]string, 0)

	for i, line := range raw {
		if i == 0 {
			// fmt.Println(line)
			for j := range line {
				// fmt.Println(j)
				// fmt.Println(line[j])
				headers = append(headers, line[j])
			}
		} else {
			rows = append(rows, line)
		}
	}

	d2file.Headers = headers
	d2file.Rows = rows
	return d2file
}

func ReadFile(filepath string) *File {
	file, err := os.Open(filepath)
	check(err)

	reader := bufio.NewReader(file)

	d2file := readCsv(reader, "filename?")

	return d2file
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
