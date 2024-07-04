package utils

import (
	"fmt"
	"os"
)

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Printf("Error closing file: %v\n", err)
	}
}
