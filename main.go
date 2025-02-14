package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string
	Completed   bool
}

func main() {
	tasks := make([]Task, 0, 20)

	fmt.Println("Please print what you want to do today (Leave blank to finish)")

	current_task_num := func() int { return len(tasks) + 1 }

	for {
		line, err := todo_usr_input(current_task_num())
		if err != nil {
			fmt.Println("Error Recieved ", err)
			os.Exit(-1)
		}

		if line == "" {
			break
		}

		tasks = append(tasks, Task{line, false})
	}

	fmt.Println("\nTask View: ")

	printCheckmark := func(task Task) string {
		if task.Completed {
			return "[x]"
		}
		return "[ ]"
	}

	for _, task := range tasks {
		fmt.Printf("%s %s\n", printCheckmark(task), task.Description)
	}

	// TODO: Reformat the json system to instead take in tasks as they go and clear them
	create_tasks_json(tasks)
}

func todo_usr_input(current_task_num int) (string, error) {
	fmt.Printf("Task %d: ", current_task_num)
	std_in := bufio.NewReader(os.Stdin)
	line, err := std_in.ReadString('\n')
	line = strings.Split(line, "\n")[0]

	if err != nil {
		return "", err
	}

	return line, nil
}

func create_tasks_json(task_list []Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(task_list); err != nil {
		return err
	}

	return nil
}
