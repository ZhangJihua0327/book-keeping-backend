package repository

import (
	"book-keeping-backend/internal/model"
)

type WorkRecordRepository struct{}

func NewWorkRecordRepository() *WorkRecordRepository {
	return &WorkRecordRepository{}
}

func (r *WorkRecordRepository) Create(record *model.WorkRecord) error {
	return DB.Create(record).Error
}

// GetByDate 查询指定日期创建的所有记录 按照主键id 倒序
func (r *WorkRecordRepository) GetByDate(dateStr string) ([]model.WorkRecord, error) {
	var records []model.WorkRecord
	err := DB.Where("date = ?", dateStr).Order("id DESC").Find(&records).Error
	return records, err
}

// Update 更新记录
func (r *WorkRecordRepository) Update(id uint64, updates map[string]interface{}) error {
	return DB.Model(&model.WorkRecord{}).Where("id = ?", id).Updates(updates).Error
}

// UpdateByRecordID 根据 record_id 更新记录 (Assuming this might be what was requested, providing as alternative or primary if needed, but PK is safer for general updates)
func (r *WorkRecordRepository) UpdateByRecordID(recordID string, updates map[string]interface{}) error {
	return DB.Model(&model.WorkRecord{}).Where("record_id = ?", recordID).Updates(updates).Error
}
