package indexer

type InvertedIndex map[string][]string

type DataFile struct {
	Path  string
	Token []string
}

func BuildInvertedIndex(files []DataFile) InvertedIndex {
	index := make(InvertedIndex) //Create an object
	for _, file := range files { //Foreach file
		seen := make(map[string]bool)      //Create a list of word added to the file
		for _, token := range file.Token { //Foreach word in the file
			if !seen[token] { //If the word doesn't already have the file path - could be removed to add a system of weight
				index[token] = append(index[token], file.Path) //Add the file path to the word
				seen[token] = true                             //Mark the word as processed
			}
		}
	}
	return index
}
