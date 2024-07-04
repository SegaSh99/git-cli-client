package main

import (
	"fmt"
	"git-cli-client/cmd"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	cmd.Execute()
}
