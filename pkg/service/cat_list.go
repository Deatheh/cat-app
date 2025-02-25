package service

import (
	"github.com/Deatheh/cat-app"
	"github.com/Deatheh/cat-app/pkg/repository"
)

type CatListService struct {
	repo repository.CatList
}

func NewCatListService(repo repository.CatList) *CatListService {
	return &CatListService{repo: repo}
}

func (s *CatListService) Create(userId int, list cat.CatList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *CatListService) GetAll(userId int) ([]cat.CatList, error) {
	return s.repo.GetAll(userId)
}

func (s *CatListService) GetById(userId, id int) (cat.CatList, error) {
	return s.repo.GetById(userId, id)
}

func (s *CatListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

func (s *CatListService) Update(userId, id int, input cat.UpdeteListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
