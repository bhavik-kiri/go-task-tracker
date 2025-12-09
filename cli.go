package main

import (
	"fmt"
	"os"
	"strconv"
)

func ParseID(raw string) int {
	id, err := strconv.Atoi(raw)
	if err != nil || id <= 0 {
		fmt.Println("Error: invalid Id: ", raw)
		os.Exit(1)
	}
	return id
}

func ParseRequiredID() int {
	if len(os.Args) < 3 {
		fmt.Println("Error: missing ID")
		os.Exit(1)
	}

	return ParseID(os.Args[2])
}

func PrintUsage() {
	fmt.Println(`Usage:
	
	add "description"
	update <id> "new description"
	delete <id>
	mark-in-progress <id>
	mark-done <id>

	list
	list todo
	list done
	list in-progress`)
}
