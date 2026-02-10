package service

import (
	"book-keeping-backend/internal/model"
	"book-keeping-backend/internal/repository"
)

type TrunkModelService struct {
	repo *repository.TrunkModelRepository
}

func NewTrunkModelService(repo *repository.TrunkModelRepository) *TrunkModelService {
	return &TrunkModelService{repo: repo}
}

func (s *TrunkModelService) AddTrunkModel(name string) error {
	tm := &model.TrunkModel{TrunkModel: name}
	return s.repo.Create(tm)
}

func (s *TrunkModelService) GetAllTrunkModels() ([]model.TrunkModel, error) {
	return s.repo.GetAll()
}
