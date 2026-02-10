package service

import (
	"book-keeping-backend/internal/model"
	"book-keeping-backend/internal/repository"
)

type WorkRecordService struct {
	repo *repository.WorkRecordRepository
}

func NewWorkRecordService(repo *repository.WorkRecordRepository) *WorkRecordService {
	return &WorkRecordService{repo: repo}
}

func (s *WorkRecordService) AddRecord(record *model.WorkRecord) error {
	return s.repo.Create(record)
}

func (s *WorkRecordService) GetRecordsByDate(dateStr string) ([]model.WorkRecord, error) {
	return s.repo.GetByDate(dateStr)
}

func (s *WorkRecordService) UpdateRecord(id uint64, updates map[string]interface{}) error {
	return s.repo.Update(id, updates)
}
