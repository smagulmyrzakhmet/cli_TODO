package handler

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/smagulmyrzakhmet/cli_TODO/internal/models"
	"github.com/smagulmyrzakhmet/cli_TODO/internal/service"
)

var command = "\033[1m=== Менеджер задач ===\033[0m\n" +
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
	reader  *bufio.Reader
}

func NewCLIHandler(service service.TaskService, reader *bufio.Reader) *CLIHandler {
	return &CLIHandler{service: service, reader: reader}
}

func (h *CLIHandler) Run() {
	for {
		fmt.Println(command)

		action, err := h.readString("Выберите действие:", 1, 2)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
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
			h.changeStatus()
		case "7":
			h.getListByStatus()
		case "0":
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Unknown action")
		}
	}
}

func (h *CLIHandler) add() {
	defer fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")
	title, err := h.readString("Название задачи", 3, 25)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	description, err := h.readString("Описание задачи", 0, 255)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	now := time.Now()
	task := models.TaskCreate{
		Title:       title,
		Description: description,
		CreateAt:    now,
	}
	_, err = h.service.Add(task)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("\033[1mЗадача успешно создана\033[0m")
}

func (h *CLIHandler) list() {
	defer fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")

	list, err := h.service.GetList()
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Список задач пустой")
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
}

func (h *CLIHandler) getById() {
	fmt.Println("----------------------------------")
	defer fmt.Println("----------------------------------")
	id, err := h.readUint("Введите id задачи")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	task, err := h.service.Get(id)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Задача не найдена")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id: %d\nНазвание: %s\nОписание: %s\nСтатус: %s\nДата создания %v\n",
		task.Id, task.Title, task.Description, task.Status, task.CreatedAt)
}

func (h *CLIHandler) update() {
	defer fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")
	id, err := h.readUint("Введите id задачи")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	task, err := h.service.Get(id)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Задача не найдена")
		return
	}
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	title, err := h.readString("Название задачи", 3, 25)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	description, err := h.readString("Описание задачи", 0, 255)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	taskUpdate := models.TaskUpdate{Title: title, Description: description}
	err = h.service.Update(task.Id, taskUpdate)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}

func (h *CLIHandler) delete() {
	defer fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")
	id, err := h.readUint("Введите id задачи")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	err = h.service.Delete(id)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Задача не найдена")
		return
	}
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}

func (h *CLIHandler) changeStatus() {
	defer fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")
	id, err := h.readUint("Введите id задачи")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	task, err := h.service.Get(id)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Задача не найдена")
		return
	}

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Текущий статус задачи %s\n", task.Status)
	fmt.Printf("\t1) %s\n2) %s\n\t3) %s\n", models.ToDo, models.InProgress, models.Done)
	option, err := h.readUint("Выберите вариант")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var status models.Status
	switch option {
	case 1:
		status = models.ToDo
	case 2:
		status = models.InProgress
	case 3:
		status = models.Done
	default:
		fmt.Println("такого варианта не существует")
		// TODO: сделать цикличный ввод до момента ввода правильного варианта или выхода
		return
	}
	err = h.service.ChangeStatus(id, status)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}

func (h *CLIHandler) getListByStatus() {
	defer fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")
	fmt.Println("Выберите статус")
	fmt.Printf("\t1) %s\n2) %s\n\t3) %s\n", models.ToDo, models.InProgress, models.Done)
	option, err := h.readUint("Выберите вариант:")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var status models.Status
	switch option {
	case 1:
		status = models.ToDo
	case 2:
		status = models.InProgress
	case 3:
		status = models.Done
	default:
		fmt.Println("такого варианта не существует")
		// TODO: сделать цикличный ввод до момента ввода правильного варианта или выхода
		return
	}
	list, err := h.service.GetListByStatus(status)
	if errors.Is(err, models.TaskNotFoundError) {
		fmt.Println("Список задач пустой")
		return
	}

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Список задач:")
	for _, task := range list {
		fmt.Println("-----------------")
		fmt.Printf("id: %d\nНазвание: %s\nСтатус: %s\n", task.Id, task.Title, task.Status)
		fmt.Println("-----------------")
	}
}

func (h *CLIHandler) readString(message string, min, max int) (string, error) {
	for {
		fmt.Println(message)

		line, err := h.reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		line = strings.TrimSpace(line)

		if len(line) >= min && len(line) <= max {
			return line, nil
		}

		fmt.Printf("Длина должна быть от %d до %d\n", min, max)
	}
}

func (h *CLIHandler) readUint(message string) (uint, error) {
	for {
		fmt.Println(message)

		line, err := h.reader.ReadString('\n')
		if err != nil {
			return 0, err
		}

		line = strings.TrimSpace(line)

		num, err := strconv.ParseUint(line, 10, 64)
		if err == nil {
			return uint(num), nil
		}

		fmt.Println("Введите корректное число")
	}
}
