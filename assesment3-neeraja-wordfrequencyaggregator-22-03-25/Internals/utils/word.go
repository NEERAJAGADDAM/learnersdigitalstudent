package utils

import (
	"fmt"
	"sync"
	"wordfrequency/Internals/models"
	"wordfrequency/Internals/services"
)

const CHUNK_SIZE = 300
const OUTPUT_FILE = "wordfrequency.csv"

func WordFrequency() {
	filename := "data.txt"
	chunks := make(chan string)
	freqChan := make(chan map[string]int)
	var wg sync.WaitGroup

	go services.ReadFileChunks(filename, CHUNK_SIZE, chunks)

	for chunk := range chunks {
		wg.Add(1)
		go services.CountFrequency(chunk, freqChan, &wg)
	}

	go func() {
		wg.Wait()
		close(freqChan)
	}()

	finalFreq := make(map[string]int)
	for freqMap := range freqChan {
		for word, count := range freqMap {
			finalFreq[word] += count
		}
	}

	for word, count := range finalFreq {
		fmt.Println(word, count)
	}

	err := models.ExportToCSV(finalFreq, OUTPUT_FILE)
	if err != nil {
		fmt.Println("Error exporting to CSV:", err)
	} else {
		fmt.Println("Word frequency data saved to", OUTPUT_FILE)
	}
}
