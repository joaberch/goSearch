package streamer

import (
	"Go-LocalSearchEngine/normalizer"
	"bufio"
	"log"
	"os"
)

func Stream(path string) []string {
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

	scanner := bufio.NewScanner(file) //Line length CAN be longer than 65536, determine max capacity at 10Mo
	//Upgrade buffer size to 1Mo
	const maxCapacity = 10 * 1024 * 1024 //10Mo
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var tokens []string
	for scanner.Scan() {
		line := scanner.Text()
		normalizedLine := normalizer.Normalize(line)
		tokens = append(tokens, normalizedLine...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(path, "Error :", err)
	}

	return tokens
}
