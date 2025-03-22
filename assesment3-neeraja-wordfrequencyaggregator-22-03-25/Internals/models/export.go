package models

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ExportToCSV(wordFreq map[string]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Word", "Frequency"})

	for word, count := range wordFreq {
		writer.Write([]string{word, strconv.Itoa(count)})
	}

	return nil
}
