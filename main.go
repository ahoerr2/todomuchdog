package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := make([]string, 0, 20)

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

		tasks = append(tasks, line)
	}

	fmt.Println("\nTask View: ")
	for _, task := range tasks {
		fmt.Println("[ ] ", task)
	}

	create_tasks_csv(tasks)
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

func create_tasks_csv(task_list []string) error {
	file, err := os.Create("tasks.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	csv_stdin := csv.NewWriter(file)

	defer csv_stdin.Flush()

	header := []string{"Task_Desc", "Completed"}
	if err := csv_stdin.Write(header); err != nil {
		return err
	}

	for _, task := range task_list {
		task_tuple := []string{task, "N"}
		if err := csv_stdin.Write(task_tuple); err != nil {
			return err
		}
	}

	return nil
}
