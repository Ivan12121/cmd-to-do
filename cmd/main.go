package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"complete"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskStore struct {
	Tasks []Task `json:"tasks"`
	nextId int
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask()
	case "list":
		lostTasks()
	case "done":
		deleteTask()
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