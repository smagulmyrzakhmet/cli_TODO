package cache

import (
	"sync/atomic"

	"github.com/smagulmyrzakhmet/cli_TODO/internal/models"
)

type RepositoryImpl struct {
	data        map[uint]models.Task
	idGenerator atomic.Uint64
}

func NewRepositoryImpl() *RepositoryImpl {
	data := map[uint]models.Task{}
	idGenerator := atomic.Uint64{}
	return &RepositoryImpl{data, idGenerator}
}

func (s *RepositoryImpl) Add(taskCreate models.TaskCreate) (models.Task, error) {
	id := uint(s.idGenerator.Add(1))
	task := models.Task{
		Id:          id,
		Title:       taskCreate.Title,
		Description: taskCreate.Description,
		CreatedAt:   taskCreate.CreateAt,
		Status:      models.InProgress,
	}
	s.data[id] = task
	return task, nil
}

func (s *RepositoryImpl) Update(id uint, taskUpdate models.TaskUpdate) error {
	task, ok := s.data[id]
	if !ok {
		return models.TaskNotFoundError
	}
	task.Title = taskUpdate.Title
	task.Description = taskUpdate.Description
	s.data[id] = task
	return nil
}

func (s *RepositoryImpl) Delete(id uint) error {
	_, ok := s.data[id]
	if !ok {
		return models.TaskNotFoundError
	}
	delete(s.data, id)
	return nil
}

func (s *RepositoryImpl) Get(id uint) (models.Task, error) {
	task, ok := s.data[id]
	if !ok {
		return models.Task{}, models.TaskNotFoundError
	}
	return task, nil
}

func (s *RepositoryImpl) GetList() ([]models.Task, error) {
	if len(s.data) == 0 {
		return nil, models.TaskNotFoundError
	}
	tasks := make([]models.Task, 0, len(s.data))
	for _, task := range s.data {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *RepositoryImpl) GetListByStatus(status models.Status) ([]models.Task, error) {
	if len(s.data) == 0 {
		return nil, models.TaskNotFoundError
	}
	tasks := make([]models.Task, 0)
	for _, task := range s.data {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	if len(tasks) == 0 {
		return nil, models.TaskNotFoundError
	}
	return tasks, nil
}

func (s *RepositoryImpl) ChangeStatus(id uint, status models.Status) error {
	task, ok := s.data[id]
	if !ok {
		return models.TaskNotFoundError
	}
	task.Status = status
	s.data[id] = task
	return nil
}
