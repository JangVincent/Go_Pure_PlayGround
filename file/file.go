package file

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)


type Task struct {
	Id int
	Task string
	Completed bool
}

type TaskManager struct {
	list []Task
	count int
}

func (taskInfo *TaskManager) RegisterTaskInFile(task string) {
	file, err := os.OpenFile("./.task", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(task + " false" + "\n")
	writer.Flush()
	taskInfo.count++
	taskInfo.list = append(taskInfo.list, Task{
		Id: taskInfo.count,
		Task: task,
		Completed: false,
	})
}

func (taskInfo *TaskManager) ReadAllTasks() ([]Task, int) {

	if taskInfo.count != 0 {
		return taskInfo.list, taskInfo.count
	}

	file, err := os.Open("./.task")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var tasks []Task
	count:=1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		taskSentence := strings.Split(scanner.Text(), " ") // id task
		task := Task{
			Id : count,
			Task : taskSentence[0],
			Completed: taskSentence[1] == "true",
		}

		tasks = append(tasks, task)
		count++
	}

	taskInfo.list = tasks
	taskInfo.count = count
	return tasks, count
}

func (taskInfo *TaskManager) MarkTaskAsCompleted(taskId int) {
	file, err := os.OpenFile("./.task", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	tasks, _ := taskInfo.ReadAllTasks()
	tasks[taskId-1].Completed = true

	file.Truncate(0)
	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		writer.WriteString(task.Task + " " + strconv.FormatBool(task.Completed) + "\n")
		writer.Flush()
	}

	taskInfo.list[taskId-1].Completed = true
}

func (taskInfo *TaskManager) DeleteTask(taskId int) {

	tasks, _ := taskInfo.ReadAllTasks()	

	for i, task := range tasks {
		if task.Id == taskId {
			tasks = slices.Delete(tasks, i, i+1)
		}
	}

	file, err := os.OpenFile("./.task", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Truncate(0)
	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		writer.WriteString(task.Task + " " + strconv.FormatBool(task.Completed) + "\n")
		writer.Flush()
	}


	taskInfo.list = slices.Delete(taskInfo.list, taskId-1, taskId)
	taskInfo.count--
}
