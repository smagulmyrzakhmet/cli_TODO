package service

import (
	"github.com/smagulmyrzakhmet/cli_TODO/internal/models"
	"github.com/smagulmyrzakhmet/cli_TODO/internal/storage"
)

type CLIServiceImpl struct {
	repo storage.Repository
}

func (s *CLIServiceImpl) Add(taskCreate models.TaskCreate) (models.Task, error) {
	return s.repo.Add(taskCreate)
}

func (s *CLIServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *CLIServiceImpl) Update(id uint, taskUpdate models.TaskUpdate) error {
	return s.repo.Update(id, taskUpdate)
}

func (s *CLIServiceImpl) Get(id uint) (models.Task, error) {
	return s.repo.Get(id)
}

func (s *CLIServiceImpl) GetList() ([]models.Task, error) {
	return s.repo.GetList()
}

func (s *CLIServiceImpl) GetListByStatus(status models.Status) ([]models.Task, error) {
	return s.repo.GetListByStatus(status)
}

func (s *CLIServiceImpl) ChangeStatus(id uint, status models.Status) error {
	return s.repo.ChangeStatus(id, status)
}

func NewCLIServiceImpl(repo storage.Repository) *CLIServiceImpl {
	return &CLIServiceImpl{repo: repo}
}
