package main

import (
	"fmt"
	"os"
)

func main() {
	query := os.Args[1:]
	fmt.Println(len(query))
}
