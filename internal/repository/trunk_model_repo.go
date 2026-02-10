package repository

import (
	"book-keeping-backend/internal/model"
)

type TrunkModelRepository struct{}

func NewTrunkModelRepository() *TrunkModelRepository {
	return &TrunkModelRepository{}
}

func (r *TrunkModelRepository) Create(trunkModel *model.TrunkModel) error {
	return DB.Create(trunkModel).Error
}

func (r *TrunkModelRepository) GetAll() ([]model.TrunkModel, error) {
	var models []model.TrunkModel
	err := DB.Find(&models).Error
	return models, err
}
