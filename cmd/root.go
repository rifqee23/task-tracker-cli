package main

import (
	"fmt"
	"github.com/rifqee23/task-tracker-cli/internal"
	"os"
	"strconv"
	"strings"
)

func main() {
	internal.Init()
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"./...\"")
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		description := strings.Join(os.Args[2:], " ")
		task := internal.Addtask(description)
		fmt.Println("Output: Task added successfully " + strconv.Itoa(task.Id))
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be integer")
			return
		}
		description := strings.Join(os.Args[3:], " ")
		task, err := internal.UpdateTask(id, description)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Updated: Task updated successfully " + strconv.Itoa(task.Id))
	case "remove":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be integer")
			return
		}
		err = internal.DeleteTask(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Remove: Task removed successfully " + strconv.Itoa(id))
	case "mark-in-progress":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be integer")
			return
		}
		_, err = internal.MarkProgress(id)
		if err != nil {
			return
		}
		fmt.Println("task marked in-progress")
	case "mark-done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID must be integer")
			return
		}
		_, err = internal.MarkDone(id)
		if err != nil {
			return
		}
		fmt.Println("task marked done")
	case "lists":
		if len(os.Args) == 2 {
			tasks := internal.GetAll()
			for _, t := range tasks {
				fmt.Printf("ID: %d | %s | %s\n", t.Id, t.Description, t.Status)
			}
		} else {
			status := os.Args[2]
			tasks, err := internal.GetByStatus(status)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, t := range tasks {
				fmt.Printf("ID: %d | %s | %s\n", t.Id, t.Description, t.Status)
			}
		}
	default:
		print("Unknown command")
	}
}
