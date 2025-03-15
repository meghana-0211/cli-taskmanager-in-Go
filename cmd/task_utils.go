package cmd

import "fmt"

type Task struct {
	ID   int
	Name string
}

var tasks = []Task{} // In-memory task list for now

// addTask adds a task to the list
func addTask(name string) {
	task := Task{ID: len(tasks) + 1, Name: name}
	tasks = append(tasks, task)
	fmt.Printf("Task added: %d - %s\n", task.ID, task.Name)
}

// listTasks prints all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Task List:")
	for _, task := range tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}
}

// removeTask removes a task by ID
func removeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Removed task: %d - %s\n", task.ID, task.Name)
			return
		}
	}
	fmt.Println("Task ID not found.")
}
