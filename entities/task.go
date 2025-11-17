package entities

import (
	"math/rand"
)

type Task struct {
	Id int64  `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
}

func NewTask() *Task {
	return &Task{
		Id: 0,
		Name: "",
		Done: false,
	}
}

func (t *Task) WithID(id int64) *Task {
	t.Id = id
	return t
}

func (t *Task) WithName(taskName string) *Task {
	t.Name = taskName
	return t
}

func (t *Task) WithDone(done bool) *Task {
	t.Done = done 
	return t
}

func NewTaskWithName(taskName string) (Task, error) {
	task := Task{
		Id:   int64(rand.Intn(10000000)),
		Name: taskName,
		Done: false,
	}

	return task, nil
}