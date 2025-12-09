package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func (list *TaskList) nextID() int {
	max := 0
	for _, t := range list.Tasks {
		if t.ID > max {
			max = t.ID
		}
	}

	return max + 1
}

func (list *TaskList) findIndex(id int) (int, error) {
	for i, t := range list.Tasks {
		if t.ID == id {
			return i, nil
		}
	}

	return -1, fmt.Errorf("task with ID %d not found", id)
}

// ---- CRUD ----

func (list *TaskList) Add(desc string) Task {
	now := time.Now().UTC()
	task := Task{
		ID:          list.nextID(),
		Description: desc,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	list.Tasks = append(list.Tasks, task)
	return task
}

func (list *TaskList) Update(id int, newDesc string) error {
	i, err := list.findIndex(id)
	if err != nil {
		return err
	}

	list.Tasks[i].Description = newDesc
	list.Tasks[i].UpdatedAt = time.Now().UTC()
	return nil
}

func (list *TaskList) Delete(id int) error {
	i, err := list.findIndex(id)
	if err != nil {
		return err
	}

	list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
	return nil
}

// ---- Status ----

func (list TaskList) SetStatus(id int, status string) error {
	valid := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	if !valid[status] {
		return errors.New("invalid status")
	}

	i, err := list.findIndex(id)
	if err != nil {
		return err
	}

	list.Tasks[i].Status = status
	list.Tasks[i].UpdatedAt = time.Now().UTC()

	return nil
}

// ---- List ----

func (list *TaskList) Print(filter string) {
	filter = strings.ToLower(filter)
	fmt.Printf("%-5s %-12s %-40s %-25s %-25s \n", "ID", "Status", "Description", "Created At", "Updated At")
	fmt.Println(strings.Repeat("-", 115))

	for _, t := range list.Tasks {
		if filter != "all" && filter != "" && t.Status != filter {
			continue
		}

		fmt.Printf("%-5d %-12s %-40s %-25s %-25s\n",
			t.ID,
			t.Status,
			trunc(t.Description, 40),
			t.CreatedAt.Format(time.RFC3339),
			t.UpdatedAt.Format(time.RFC3339),
		)
	}
}

func trunc(s string, n int) string {
	if len(s) <= n {
		return s
	}

	return s[:n-3] + "..."
}

// func (todos *Todos) add(title string) {
// 	todo := Todo{
// 		Title:       title,
// 		Completed:   false,
// 		CompletedAt: nil,
// 		CreatedAt:   time.Now(),
// 	}

// 	*todos = append(*todos, todo)
// }

// func (todos *Todos) validateIndex(index int) error {
// 	if index < 0 || index >= len(*todos) {
// 		err := errors.New("invalid index")
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }

// func (todos *Todos) delete(index int) error {
// 	t := *todos

// 	if err := t.validateIndex(index); err != nil {
// 		return err
// 	}

// 	*todos = append(t[:index], t[index+1:]...)
// 	return nil
// }

// func (todos *Todos) toggle(index int) error {
// 	t := *todos

// 	if err := t.validateIndex(index); err != nil {
// 		return err
// 	}

// 	isCompleted := t[index].Completed

// 	if !isCompleted {
// 		completionTime := time.Now()
// 		t[index].CompletedAt = &completionTime
// 	}

// 	t[index].Completed = !isCompleted

// 	return nil
// }

// func (todos *Todos) edit(index int, title string) error {
// 	t := *todos

// 	if err := t.validateIndex(index); err != nil {
// 		return err
// 	}

// 	t[index].Title = title

// 	return nil
// }

// func (todos *Todos) print() {

// 	table := table.New(os.Stdout)
// 	table.SetRowLines(false)
// 	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
// 	for index, t := range *todos {
// 		completed := "❌"
// 		completedAt := ""

// 		if t.Completed {
// 			completed = "✅"

// 			if t.CompletedAt != nil {
// 				completedAt = t.CompletedAt.Format(time.RFC1123)
// 			}
// 		}

// 		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)

// 	}

// 	table.Render()
// }
