package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID      int
	Title   string
	Complete bool
}

var tasks []Task
var taskID int

func main() {
	fmt.Println("Welcome to the To-Do List App!")
	fmt.Println("Commands: add <task>, list, complete <id>, delete <id>, exit")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <task>")
				continue
			}
			addTask(parts[1])
		case "list":
			listTasks()
		case "complete":
			if len(parts) < 2 {
				fmt.Println("Usage: complete <id>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid task ID.")
				continue
			}
			completeTask(id)
		case "delete":
			if len(parts) < 2 {
				fmt.Println("Usage: delete <id>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid task ID.")
				continue
			}
			deleteTask(id)
		default:
			fmt.Println("Unknown command. Available commands: add, list, complete, delete, exit")
		}
	}
}

func addTask(title string) {
	taskID++
	tasks = append(tasks, Task{ID: taskID, Title: title, Complete: false})
	fmt.Printf("Added task: %d - %s\n", taskID, title)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := "Incomplete"
		if task.Complete {
			status = "Complete"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}

func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Complete = true
			fmt.Printf("Task %d marked as complete.\n", id)
			return
		}
	}
	fmt.Printf("Task %d not found.\n", id)
}

func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task %d deleted.\n", id)
			return
		}
	}
	fmt.Printf("Task %d not found.\n", id)
}
