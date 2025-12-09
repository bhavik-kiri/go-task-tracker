# Task Tracker CLI

A simple **command-line task tracker** built in **Go** to help you manage your tasks.  
You can **add, update, delete, mark as in-progress or done, and list tasks**. All tasks are stored in a JSON file in the current directory.

---

## Features

- Add new tasks  
- Update existing tasks  
- Delete tasks  
- Mark tasks as **in-progress** or **done**  
- List tasks by status: `todo`, `in-progress`, `done`, or all  
- Tasks stored in a JSON file (`todos.json`)  
- CLI uses **positional arguments**, no external libraries required  

---

## Task Structure

Each task has the following properties:

| Field        | Type   | Description                       |
| ------------ | ------ | --------------------------------- |
| `id`        | int    | Unique identifier                 |
| `description` | string | Task description                  |
| `status`    | string | Task status (`todo`, `in-progress`, `done`) |
| `createdAt` | string | Timestamp when task was created   |
| `updatedAt` | string | Timestamp when task was last updated |

---

## Usage

Run the program using `go run ./` followed by the command and arguments:

### Add a task
```bash
go run ./ add "Buy groceries"
```

### Update a task
```bash
go run ./ update 1 "Buy groceries and cook dinner"
```

### Delete a task
```bash
go run ./ delete 1
```

### Mark task as in-progress
```bash
go run ./ mark-in-progress 1
```

### Mark task as done
```bash
go run ./ mark-done 1
```

### List tasks
```bash
go run ./ list          # List all tasks
go run ./ list todo     # List tasks with status 'todo'
go run ./ list in-progress # List tasks with status 'in-progress'
go run ./ list done     # List tasks with status 'done'
```

### Example
```bash
go run ./ add "Finish Go project"
go run ./ add "Buy Milk"
go run ./ list
go run ./ mark-in-progress 1
go run ./ mark-done 2
go run ./ list done
go run ./ update 1 "Finish Go CLI project"
go run ./ delete 2
```

## JSON Storage

Tasks are stored in todos.json in the current directory.
Example structure:
```bash
{
  "tasks": [
    {
      "id": 1,
      "description": "Finish Go CLI project",
      "status": "in-progress",
      "createdAt": "2025-12-09T15:00:00Z",
      "updatedAt": "2025-12-09T16:00:00Z"
    }
  ]
}

```

## Requirements

- Go 1.20+ (or any recent Go version)

- No external libraries required