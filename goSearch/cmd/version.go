package cmd

import "fmt"

func ShowVersion() {
	fmt.Printf("Version: %s\nDate: %s\nGo: %s\n", "0.1.0", "   28.08.2025", "     go1.25.0 windows/amd64")
}
