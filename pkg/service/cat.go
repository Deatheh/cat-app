package service

import (
	"github.com/Deatheh/cat-app"
	"github.com/Deatheh/cat-app/pkg/repository"
)

type CatService struct {
	repo     repository.Cat
	listRepo repository.CatList
}

func NewCatService(repo repository.Cat, listRepo repository.CatList) *CatService {
	return &CatService{repo: repo, listRepo: listRepo}
}

func (s *CatService) Create(userId, listId int, item cat.Cat) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *CatService) GetAll(userId, listId int) ([]cat.Cat, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *CatService) GetById(userId, listId, itemId int) (cat.Cat, error) {
	return s.repo.GetById(userId, listId, itemId)
}

func (s *CatService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *CatService) Update(userId, listId, itemId int, input cat.UpdateCatInput) error {
	return s.repo.Update(userId, listId, itemId, input)
}
