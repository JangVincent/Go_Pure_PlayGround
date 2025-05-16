package task

import (
	"fmt"
	"main/src/file"
)

var taskInfo file.TaskManager

func AddTask(task string) {
	taskInfo.RegisterTaskInFile(task)
}

func ViewAllTasks() ([]file.Task, int) {
	return taskInfo.ReadAllTasks()
}

func MarkTaskAsCompleted(taskId int) {
	taskInfo.MarkTaskAsCompleted(taskId)
}

func DeleteTask(taskId int) {
	fmt.Println("Hello, World!")
}


