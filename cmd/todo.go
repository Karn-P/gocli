package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Manage your to-do list",
	Long: `Manage your to-do list with commands to add, remove, and view tasks.
	For example:
	gocli todo add "Buy groceries"
	gocli todo remove 1
	gocli todo list
	gocli todo complete 1
	`,
	Run: todoList,
}

func readTasksFromCSV(taskFile string, cmd *cobra.Command) []Task {
	tasks := []Task{}

	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		file, err := os.Create(taskFile)
		if err != nil {
			cmd.PrintErr("Error creating todo file:", err)
			return tasks
		}
		writer := csv.NewWriter(file)
		writer.Write([]string{"id", "description", "createdAt", "iscomplete"})
		writer.Flush()
		file.Close()
	}

	file, err := os.Open(taskFile)
	if err != nil {
		cmd.PrintErr("Error opening todo file:", err)
		return tasks
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		cmd.PrintErr("Error reading tasks:", err)
		return tasks
	}

	if len(records) > 1 {
		for i := 1; i < len(records); i++ {
			record := records[i]
			if len(record) == 4 {
				id, _ := strconv.Atoi(record[0])
				createdAt, _ := time.Parse(time.RFC3339, record[2])
				isComplete, _ := strconv.ParseBool(record[3])

				task := Task{
					ID:          id,
					Description: record[1],
					CreatedAt:   createdAt,
					IsComplete:  isComplete,
				}
				tasks = append(tasks, task)
			}
		}
	}

	return tasks
}

func completeTaskFunc(taskID int, cmd *cobra.Command) {
	taskFile := "todo.csv"
	tasks := readTasksFromCSV(taskFile, cmd)

	found := false

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].IsComplete = true
			found = true
			break
		}
	}

	if !found {
		cmd.Println("Task ID not found:", taskID)
		return
	}

	if err := saveTasksToCSV(taskFile, tasks); err != nil {
		cmd.PrintErr("Error saving tasks:", err)
		return
	}
	cmd.Println("Marked task", taskID, "as complete")
}

func removeTaskFunc(taskID int, cmd *cobra.Command) {
	taskFile := "todo.csv"
	tasks := readTasksFromCSV(taskFile, cmd)

	found := false
	newTasks := []Task{}

	for _, task := range tasks {
		if task.ID != taskID {
			newTasks = append(newTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		cmd.Println("Task ID not found:", taskID)
		return
	}

	tasks = newTasks
	if err := saveTasksToCSV(taskFile, tasks); err != nil {
		cmd.PrintErr("Error saving tasks:", err)
		return
	}
	cmd.Println("Removed task with ID:", taskID)
}

func listTasksFunc(cmd *cobra.Command) {
	taskFile := "todo.csv"
	tasks := readTasksFromCSV(taskFile, cmd)

	if len(tasks) == 0 {
		cmd.Println("No tasks found.")
		return
	}

	cmd.Println("ID | Description | Created At | Completed")
	cmd.Println("---|-------------|------------|----------")

	for _, task := range tasks {
		status := "No"
		if task.IsComplete {
			status = "Yes"
		}
		cmd.Printf("%2d | %-20s | %s | %s\n",
			task.ID,
			task.Description,
			task.CreatedAt.Format("2006-01-02 15:04"),
			status)
	}
}

func addTaskFunc(description string, cmd *cobra.Command) {
	taskFile := "todo.csv"
	tasks := readTasksFromCSV(taskFile, cmd)

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          newID,
		Description: description,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}
	tasks = append(tasks, newTask)

	if err := saveTasksToCSV(taskFile, tasks); err != nil {
		cmd.PrintErr("Error saving tasks:", err)
		return
	}
	cmd.Println("Added task:", description)
}

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task to the to-do list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addTaskFunc(args[0], cmd)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		listTasksFunc(cmd)
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove [id]",
	Short: "Remove a task by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.PrintErr("Invalid ID:", err)
			return
		}
		removeTaskFunc(id, cmd)
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.PrintErr("Invalid ID:", err)
			return
		}
		completeTaskFunc(id, cmd)
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

	todoCmd.AddCommand(addCmd)
	todoCmd.AddCommand(listCmd)
	todoCmd.AddCommand(removeCmd)
	todoCmd.AddCommand(completeCmd)

	todoCmd.Flags().StringP("add", "a", "", "Add a new task to the to-do list")
	todoCmd.Flags().IntP("remove", "r", 0, "Remove a task by its number from the to-do list")
	todoCmd.Flags().BoolP("list", "l", false, "List all tasks in the to-do list")
	todoCmd.Flags().IntP("complete", "c", 0, "Mark task with ID as complete")
}

func todoList(cmd *cobra.Command, args []string) {
	addTask, _ := cmd.Flags().GetString("add")
	removeTask, _ := cmd.Flags().GetInt("remove")
	listTasks, _ := cmd.Flags().GetBool("list")
	completeTask, _ := cmd.Flags().GetInt("complete")

	if addTask != "" {
		addTaskFunc(addTask, cmd)
	} else if removeTask > 0 {
		removeTaskFunc(removeTask, cmd)
	} else if completeTask > 0 {
		completeTaskFunc(completeTask, cmd)
	} else if listTasks || (len(args) == 0 && addTask == "" && removeTask <= 0 && completeTask <= 0) {
		listTasksFunc(cmd)
	} else {
		cmd.Println("No valid command provided. Use --help for usage information.")
	}
}

func saveTasksToCSV(filename string, tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"id", "description", "createdAt", "iscomplete"}); err != nil {
		return fmt.Errorf("error writing header: %w", err)
	}

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(task.IsComplete),
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing task: %w", err)
		}
	}

	return nil
}
