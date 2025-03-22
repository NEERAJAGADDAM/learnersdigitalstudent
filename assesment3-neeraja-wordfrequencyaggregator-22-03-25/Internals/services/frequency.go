package services

import (
	"sync"
	"wordfrequency/Internals/models"
)

func CountFrequency(text string, freqChan chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	wordCounts := make(map[string]int)

	words := models.SplitWords(text)
	for _, word := range words {
		wordCounts[word]++
	}

	freqChan <- wordCounts
}
