package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Task struct {
	Description string
	Completed   bool
}

func main() {
	tasks, err := read_tasks_json()
	if err != nil {
		log.Fatal(err)
	}

	current_task_num := func() int { return len(tasks) + 1 }

	fmt.Printf("Please print what you want to do today (Leave blank to finish) (Current Tasks: %d)\n", len(tasks))

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
	update_tasks_json(tasks)
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

func read_tasks_json() ([]Task, error) {
	file, err := retrieve_tasks_file(os.O_CREATE | os.O_RDONLY)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var task_list []Task

	err = json.NewDecoder(file).Decode(&task_list)

	if err == io.EOF {
		return make([]Task, 0, 20), nil
	}

	if err != nil {
		return nil, err
	}

	return task_list, err
}

func update_tasks_json(task_list []Task) error {
	file, err := retrieve_tasks_file(os.O_TRUNC | os.O_RDWR)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(task_list); err != nil {
		return err
	}

	return nil
}

func retrieve_tasks_file(flag int) (*os.File, error) {
	file, err := os.OpenFile("tasks.json", flag, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
