package main

import (
	"fmt"
	"main/src/task"
)

func main() {
	fmt.Println("Hello Vincent Task Manager App")
	fmt.Println("=========================")
	fmt.Println("1. Add a new task")
	fmt.Println("2. View all tasks")
	fmt.Println("3. Mark a task as completed")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Exit")
	fmt.Println("=========================")

	for {
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Exiting the program...")
			break;
		}
	
		switch choice {
			case 1:
				fmt.Print("Enter a new task : ")
	
				var newTask string
				fmt.Scanln(&newTask)
				task.AddTask(newTask)
			case 2:
				tasks, _ := task.ViewAllTasks()
				for index, task := range tasks {
					fmt.Println(index+1, ". ", task.Task)
				}
			case 3:
				fmt.Print("Enter the task number to mark as completed: ")
	
				var taskId int
				fmt.Scanln(&taskId)
				task.MarkTaskAsCompleted(taskId)
			case 4:
				fmt.Print("Enter the task number for deletion: ")
	
				var taskId int
				fmt.Scanln(&taskId)
				task.DeleteTask(taskId)
			default:
				fmt.Println("Invalid Option Number");
		}

	}
	
}
