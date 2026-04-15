package storage

import "github.com/smagulmyrzakhmet/cli_TODO/internal/models"

type Repository interface {
	Add(task models.TaskCreate) (models.Task, error)
	Delete(id uint) error
	Update(id uint, task models.TaskUpdate) error
	Get(id uint) (models.Task, error)
	GetList() ([]models.Task, error)
	GetListByStatus(status models.Status) ([]models.Task, error)
	ChangeStatus(id uint, status models.Status) error
}
