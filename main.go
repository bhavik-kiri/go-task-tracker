package main

func main() {
	// fmt.Println("Inside To Do app")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
	// todos.add("Buy Milk")
	// todos.add("Buy Bread")
	// fmt.Printf("%v\n\n", todos)
	// todos.delete(0)
	// fmt.Printf("%v\n\n", todos)

	// todos.toggle(0)
	// todos.print()
	// storage.Save(todos)
}
