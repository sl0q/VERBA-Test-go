package main

import (
	"fmt"
	"os"

	"VERBA-Test/database"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: VERBA-Test [databaseCredentialsFilePath.json]")
		return
	}
	databaseCredentialsFilePath := os.Args[1]

	_, err := database.ConnectDB(databaseCredentialsFilePath)

	if err != nil {
		panic("Failed to connect to Database")
	}
}
