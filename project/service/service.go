package service

import (
	"github.com/apostrohedottilde/indymorning/api/database"
	"github.com/apostrohedottilde/indymorning/api/project/domain"
)

type Service interface {
	Create(loggedInUser string, model domain.GameProject) (domain.GameProject, error)
	FindOne(loggedInUser string, id string) (domain.GameProject, error)
	FindAll(loggedInUser string) ([]domain.GameProject, error)
	Update(loggedInUser string, id string, update domain.GameProject) (domain.GameProject, error)
	Delete(loggedInUser string, s string) error
}

type ProjectService struct {
	repository *database.ProjectRepository
}

func (ps *ProjectService) Create(loggedInUser string, model domain.GameProject) (domain.GameProject, error) {
	model.Creator = loggedInUser
	dummy, err := ps.repository.Create(model)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) FindOne(loggedInUser string, id string) (domain.GameProject, error) {
	dummy, err := ps.repository.FindOne(id)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) FindAll(loggedInUser string) ([]domain.GameProject, error) {
	dummy, err := ps.repository.FindAll()
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) Update(loggedInUser string, id string, update domain.GameProject) (domain.GameProject, error) {
	dummy, err := ps.repository.Update(id, update)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

func (ps *ProjectService) Delete(loggedInUser string, id string) error {
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
