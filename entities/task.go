package entities

import (
	"math/rand"
)

type Task struct {
	Id int64  `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
}

func NewTask() Task {
	return Task{
		Id: 0,
		Name: "",
		Done: false,
	}
}

func NewTaskWithName(taskName string) (Task, error) {
	task := Task{
		Id:   int64(rand.Intn(10000000)),
		Name: taskName,
		Done: false,
	}

	return task, nil
}