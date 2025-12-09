package main

import (
	"fmt"
	"os"
)

const jsonFile = "todos.json"

func main() {
	if len(os.Args) < 2 {

		PrintUsage()
		return
	}

	command := os.Args[1]

	store := NewStorage[TaskList](jsonFile)

	var tasks TaskList
	store.Load(&tasks)

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task description")
			PrintUsage()
			return
		}

		desc := os.Args[2]
		t := tasks.Add(desc)
		store.Save(tasks)
		fmt.Printf("Task added successfully (ID: %d)\n", t.ID)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: Updare <id> \"new description\"")
			PrintUsage()
			return
		}

		id := ParseID(os.Args[2])
		err := tasks.Update(id, os.Args[3])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		store.Save(tasks)
		fmt.Println("Task updated successfully")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			PrintUsage()
			return
		}
		id := ParseID(os.Args[2])
		err := tasks.Delete(id)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		store.Save(tasks)
		fmt.Println("Task Deleted")

	case "mark-in-progress":
		id := ParseRequiredID()
		err := tasks.SetStatus(id, "in-progress")
		if err != nil {
			fmt.Println("Error: ", err)
			PrintUsage()
			return
		}

		store.Save(tasks)
		fmt.Println("Task marked as in-progress")

	case "mark-done":
		id := ParseRequiredID()
		err := tasks.SetStatus(id, "done")
		if err != nil {
			fmt.Println("Error: ", err)
			PrintUsage()
			return
		}

		store.Save(tasks)
		fmt.Println("Task marked as done")

	case "list":
		status := "all"

		if len(os.Args) > 2 {
			status = os.Args[2]
		}

		fmt.Println(status)
		tasks.Print(status)

	default:
		fmt.Println("Unknown command: ", command)
		PrintUsage()
	}
}
