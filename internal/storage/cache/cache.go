package cache

import (
	"sync/atomic"

	"github.com/smagulmyrzakhmet/cli_TODO/internal/models"
)

type CacheRepositoryImpl struct {
	data        map[uint]models.Task
	idGenerator atomic.Uint64
}

func NewCacheRepositoryImpl() *CacheRepositoryImpl {
	data := map[uint]models.Task{}
	idGenerator := atomic.Uint64{}
	return &CacheRepositoryImpl{data, idGenerator}
}

func (s *CacheRepositoryImpl) Add(taskCreate models.TaskCreate) (models.Task, error) {
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

func (s *CacheRepositoryImpl) Update(id uint, taskUpdate models.TaskUpdate) error {
	task, ok := s.data[id]
	if !ok {
		return models.TaskNotFoundError
	}
	task.Title = taskUpdate.Title
	task.Description = taskUpdate.Description
	s.data[id] = task
	return nil
}

func (s *CacheRepositoryImpl) Delete(id uint) error {
	_, ok := s.data[id]
	if !ok {
		return models.TaskNotFoundError
	}
	delete(s.data, id)
	return nil
}

func (s *CacheRepositoryImpl) Get(id uint) (models.Task, error) {
	task, ok := s.data[id]
	if !ok {
		return models.Task{}, models.TaskNotFoundError
	}
	return task, nil
}

func (s *CacheRepositoryImpl) GetAll() ([]models.Task, error) {
	if len(s.data) == 0 {
		return nil, models.TaskNotFoundError
	}
	tasks := make([]models.Task, 0, len(s.data))
	for _, task := range s.data {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *CacheRepositoryImpl) GetListByStatus(status models.Status) ([]models.Task, error) {
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

func (s *CacheRepositoryImpl) ChangeStatus(id uint, status models.Status) error {
	task, ok := s.data[id]
	if !ok {
		return models.TaskNotFoundError
	}
	task.Status = status
	s.data[id] = task
	return nil
}
