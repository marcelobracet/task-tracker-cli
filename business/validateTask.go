package business

import (
	"errors"
	"fmt"
	"task-tracker/entities"
)

func ValidateTaskName(taskName string) error {
	if taskName == "" {
		return errors.New("taskName should not be empty")
	}
	return nil
}

func ValidateTask(task entities.Task) entities.Task {
	if(task.Name == ""){
		fmt.Println("task should not be and empty string")
	}

	if(task.Id == 0){
		fmt.Println("ID should be greater then 0")
	}

	return task
}