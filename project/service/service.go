package service

import (
	"github.com/apostrohedottilde/indymorning/api/database"
	"github.com/apostrohedottilde/indymorning/api/project/domain"
)

type Service interface {
	Create(model domain.GameProject) (domain.GameProject, error)
	FindOne(id string) (domain.GameProject, error)
	FindAll() ([]domain.GameProject, error)
	Update(id string, update domain.GameProject) (domain.GameProject, error)
	Delete(s string) error
}

type ProjectService struct {
	repository *database.ProjectRepository
}

func (ps *ProjectService) Create(model domain.GameProject) (domain.GameProject, error) {
	dummy, err := ps.repository.Create(model)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) FindOne(id string) (domain.GameProject, error) {
	dummy, err := ps.repository.FindOne(id)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) FindAll() ([]domain.GameProject, error) {
	dummy, err := ps.repository.FindAll()
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) Update(id string, update domain.GameProject) (domain.GameProject, error) {
	dummy, err := ps.repository.Update(id, update)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) Delete(id string) error {
	err := ps.repository.Delete(id)
	if err != nil {
		panic(err)
	}
	return err
}

func New(repo *database.ProjectRepository) *ProjectService {
	return &ProjectService{
		repository: repo,
	}
}
