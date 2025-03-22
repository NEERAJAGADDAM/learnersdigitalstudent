package services

import (
	"bufio"
	"os"
)

func ReadFileChunks(filename string, chunkSize int, chunks chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		close(chunks)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, chunkSize)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
			chunks <- string(buf[:n])
		}
		if err != nil {
			break
		}
	}
	close(chunks)
}
