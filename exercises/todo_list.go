package main

import (
	"fmt"
	"strings"
)

type Task struct {
	Name     string
	Complete bool
	SubTasks []*Task
}

// AddSubTask добавляет подзадачу к задаче
func (t *Task) AddSubTask(subTaskName string) {
	t.SubTasks = append(t.SubTasks, &Task{Name: subTaskName, Complete: false})
}

// MarkComplete отмечает задачу как выполненную
func (t *Task) MarkComplete() {
	t.Complete = true
	for _, subTask := range t.SubTasks {
		subTask.MarkComplete() // Рекурсивно отмечаем все подзадачи как выполненные
	}
}

// String рекурсивно форматирует задачу и ее подзадачи
func (t *Task) String() string {
	var sb strings.Builder
	if t.Complete {
		sb.WriteString("[Выполнено] ")
	} else {
		sb.WriteString("[Невыполнено] ")
	}
	sb.WriteString(t.Name)
	for _, subTask := range t.SubTasks {
		sb.WriteString("\n\t")
		sb.WriteString(subTask.String()) // Рекурсивный вызов String для подзадач
	}
	return sb.String()
}

// Задание: Объясните, как будет изменяться состояние объекта task после выполнения каждой операции в функции main.
func hardTask() {
	task := Task{Name: "Основная задача"}
	task.AddSubTask("Подзадача 1")
	task.AddSubTask("Подзадача 2")

	fmt.Println("Состояние задачи до выполнения:")
	fmt.Println(task.String())

	task.MarkComplete()

	fmt.Println("\nСостояние задачи после выполнения:")
	fmt.Println(task.String())
}
