package cmd

import "fmt"

func ShowVersion() {
	fmt.Printf("Version: %s\nDate: %s\nGo: %s\n", "1.0.0", "   01.09.2025", "     go1.25.0 windows/amd64")
}
