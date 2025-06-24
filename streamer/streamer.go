package streamer

import (
	"Go-LocalSearchEngine/normalizer"
	"bufio"
	"log"
	"os"
)

func Stream(path string) []string {
	//Stream
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file) //Line length CAN be longer than 65536, TODO - see if the Buffer Method is required
	var tokens []string
	for scanner.Scan() {
		line := scanner.Text()
		normalizedLine := normalizer.Normalize(line)
		tokens = append(tokens, normalizedLine...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tokens
}
