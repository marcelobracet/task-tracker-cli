package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"task-tracker/handlers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("┌─────────────────────────────────┐")
		fmt.Println("│          MENU DE OPÇÕES         │")
		fmt.Println("├─────────────────────────────────┤")
		fmt.Println("│  0 - Sair                       │")
		fmt.Println("│  1 - Adicionar tarefa           │")
		fmt.Println("│  2 - Atualizar tarefa           │")
		fmt.Println("│  3 - Deletar tarefa             │")
		fmt.Println("│  4 - Listar todas as tarefas    │")
		fmt.Println("│  5 - Listar tarefas não feitas  │")
		fmt.Println("└─────────────────────────────────┘")
		fmt.Println()

		fmt.Print("Escolha a opção desejada: ")
		scanner.Scan()
		userInput := scanner.Text()

		switch userInput {
		case "0":
			fmt.Println("Saindo...")
			return
		case "1":
			fmt.Print("\nQual tarefa você deseja adicionar? ")
			scanner.Scan()
			taskName := scanner.Text()
			handlers.AddTask(taskName)
		case "2":
			fmt.Print("\nQual o ID da tarefa que você deseja atualizar? ")
			scanner.Scan()
			taskId := scanner.Text()
			taskIdNumber, err := strconv.ParseInt(taskId, 10, 64)
			if err != nil {
				fmt.Println("Error converting string to int64")
			}
			fmt.Print("\nQual o nome atualizado da tarefa? ")
			scanner.Scan()
			taskName := scanner.Text()
			fmt.Print("\nA tarefa foi concluída? (s/n) ")
			scanner.Scan()
			done := scanner.Text() == "s"
			task, err := handlers.UpdateTask(taskIdNumber, taskName, done)
			if err != nil {
				fmt.Println("Error updating task:", err)
			}
			fmt.Println("Task updated successfully:", *task)
		case "3":
			fmt.Print("\nQual o ID da tarefa que você deseja deletar? ")
			scanner.Scan()
			taskId := scanner.Text()
			taskIdNumber, err := strconv.ParseInt(taskId, 10, 64)
			if err != nil {
				fmt.Println("Error converting string to int64")
				break
			}
			deletedTask, err := handlers.DeleteTask(taskIdNumber)
			if err != nil {
				fmt.Println("Error deleting task:", err)
				break
			}
			fmt.Println("Task deleted successfully:", *deletedTask)
		case "4":
			tasks, err := handlers.GetTasks()
			if err != nil {
				fmt.Println("Error getting tasks:", err)
			}
			for _, task := range tasks {
				fmt.Printf("ID: %d, Name: %s, Done: %v\n", task.Id, task.Name, task.Done)
			}
		case "5":
			tasks, err := handlers.PrintTasksNotDone()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(tasks)
		default:
			fmt.Println()
			fmt.Println("================================================")
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
			fmt.Println("================================================")
			fmt.Println()
		}
	}

}
