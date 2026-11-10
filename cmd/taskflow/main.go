package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: taskflow <command>")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		fmt.Println("Add task selected")
	case "list":
		fmt.Println("List tasks selected")
	case "complete":
		fmt.Println("Complete task selected")
	case "delete":
		fmt.Println("Delete task selected")
	default:
		fmt.Println("Unknown command", command)
	}
}
