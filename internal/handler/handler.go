package handler

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/smagulmyrzakhmet/cli_TODO/internal/models"
	"github.com/smagulmyrzakhmet/cli_TODO/internal/service"
)

var command = "=== Менеджер задач ===\n" +
	"1. Добавить задачу\n" +
	"2. Список задач\n" +
	"3. Получить задачу по ID\n" +
	"4. Обновить задачу\n" +
	"5. Удалить задачу\n" +
	"6. Изменить статус задачи\n" +
	"7. Получить список задач по статусу\n" +
	"0. Выход"

type CLIHandler struct {
	service service.TaskService
}

func (h *CLIHandler) Run() {
	for {
		fmt.Println(command)
		var action string
		_, err := fmt.Scanln(&action)
		if err != nil {
			log.Fatal(err)
		}
		switch action {
		case "1":
			h.add()
		case "2":
			h.list()
		case "3":
			h.getById()
		case "4":
			h.update()
		case "5":
			h.delete()
		case "6":
		case "7":
		case "0":
		default:
			fmt.Println("Unknown action")
		}
	}
}

func (h *CLIHandler) add() {
	fmt.Println("----------------------------------")
	title, err := scanln("Название задачи", 3, 25)
	if err != nil {
		log.Fatal(err)
	}

	description, err := scanln("Описание задачи", 0, 255)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	task := models.TaskCreate{
		Title:       title,
		Description: description,
		CreateAt:    now,
	}
	_, err = h.service.Add(task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------------------------------")
}

func (h *CLIHandler) list() {
	fmt.Println("----------------------------------")
	list, err := h.service.GetList()
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Список задач пустой")
		fmt.Println("----------------------------------")
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Список задач:")
	for _, task := range list {
		fmt.Println("-----------------")
		fmt.Printf("id: %d\nНазвание: %s\nСтатус: %s\n", task.Id, task.Title, task.Status)
		fmt.Println("-----------------")
	}
	fmt.Println("----------------------------------")
}

func (h *CLIHandler) getById() {
	fmt.Println("----------------------------------")
	fmt.Println("Введите id задачи")
	var id uint
	_, err := fmt.Scanln(&id)
	if err != nil {
		log.Fatal(err)
	}
	task, err := h.service.Get(id)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Задача не найдена")
		fmt.Println("----------------------------------")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id: %d\nНазвание: %s\nОписание: %s\nСтатус: %s\nДата создания %v\n",
		task.Id, task.Title, task.Description, task.Status, task.CreatedAt)
	fmt.Println("----------------------------------")
}

func (h *CLIHandler) update() {
	fmt.Println("----------------------------------")
	fmt.Println("Введите id задачи")
	var id uint
	_, err := fmt.Scanln(&id)
	if err != nil {
		log.Fatal(err)
	}
	task, err := h.service.Get(id)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Задача не найдена")
		fmt.Println("----------------------------------")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	title, err := scanln("Название задачи", 3, 25)
	if err != nil {
		log.Fatal(err)
	}

	description, err := scanln("Описание задачи", 0, 255)
	if err != nil {
		log.Fatal(err)
	}
	taskUpdate := models.TaskUpdate{Title: title, Description: description}
	err = h.service.Update(task.Id, taskUpdate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------------------------------")
}

func (h *CLIHandler) delete() {
	fmt.Println("----------------------------------")
	fmt.Println("Введите id задачи")
	var id uint
	_, err := fmt.Scanln(&id)
	if err != nil {
		log.Fatal(err)
	}
	err = h.service.Delete(id)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Задача не найдена")
		fmt.Println("----------------------------------")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------------------------------")
}

func (h *CLIHandler) changeStatus() {
	fmt.Println("----------------------------------")
	fmt.Println("Введите id задачи")
	var id uint
	_, err := fmt.Scanln(&id)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: доделать метод
	scanln("На какой статус поменять?", 4, 11)
	fmt.Println("----------------------------------")
}

func scanln(message string, min, max int) (string, error) {
	var line string
	for {
		fmt.Println(message)
		_, err := fmt.Scanln(&line)
		if err != nil {
			return "", err
		}
		if len(line) >= min && len(line) <= max {
			return line, nil
		}
		fmt.Println("Некорректный ввод!")
		fmt.Printf("Количество символов должно быть в диапазоне от %d до %d\n", min, max)
	}
}
