package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"task-tracker/entities"
)

const (
	FILE_NAME = "tarefas.json"
)

func AddTask(taskName string) (*entities.Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return &entities.Task{}, err
	}

	newId := int64(rand.Intn(10000000))

	task := entities.Task{
		Id:   newId,
		Name: taskName,
		Done: false,
	}

	tasks = append(tasks, task)

	if err := saveTasks(tasks); err != nil {
		return nil, err
	}

	return &task, nil
}

func saveTasks(tasks []entities.Task) error {
	file, err := os.OpenFile(FILE_NAME, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		panic(err)
	}

	return nil
}

func GetTasks() ([]entities.Task, error) {

	file, err := os.ReadFile(FILE_NAME)
	if err != nil {
		if os.IsNotExist(err) {
			return []entities.Task{}, nil
		}
		return nil, err
	}

	var tasks []entities.Task

	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func UpdateTask(taskId int64, taskName string, done bool) (*entities.Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return nil, err
	}

	var updatedTask *entities.Task

	for i := range tasks {
		task := &tasks[i]
		if task.Id == taskId {
			task.Name = taskName
			task.Done = done
			updatedTask = task
			break
		}
	}

	if updatedTask == nil {
		return nil, fmt.Errorf("task with id %d not found", taskId)
	}

	if err := saveTasks(tasks); err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func DeleteTask(taskId int64) (*entities.Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return nil, err
	}

	var deletedTask *entities.Task

	for i := range tasks {
		if tasks[i].Id == taskId {
			deletedTask = &tasks[i]
			// linha abaixo basicamente é a forma de remover um item do indice i em golang
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if deletedTask == nil {
		return nil, fmt.Errorf("task with id %d not found", taskId)
	}

	if err := saveTasks(tasks); err != nil {
		return nil, err
	}

	return deletedTask, nil
}

func PrintTasksNotDone() ([]entities.Task, error) {
	var undoneTasks []entities.Task

	tasks, err := GetTasks()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if !t.Done {
			undoneTasks = append(undoneTasks, t)
		}
	}

	return undoneTasks, nil
}

