package main

import (
	"fmt"
	"os"
	"todo-cli/internal"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		internal.AddTask()
	case "list":
		internal.ListTasks()
	case "done":
		internal.DeleteTask()
	default:
		printUsage()
	}
}

func printUsage() {
    fmt.Println("Использование:")
    fmt.Println("  todo add <название задачи>")
    fmt.Println("  todo list")
    fmt.Println("  todo done <id>")
    fmt.Println("  todo delete <id>")
}