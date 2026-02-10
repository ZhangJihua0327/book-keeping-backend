package model

import "time"

type WorkRecord struct {
	ID               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	RecordID         string    `gorm:"type:varchar(256)" json:"record_id"`
	TrunkModel       string    `gorm:"type:varchar(256)" json:"trunk_model" binding:"required"`
	Date             time.Time `gorm:"type:date" json:"date" binding:"required"`
	CustomerName     string    `gorm:"type:varchar(256)" json:"customer_name" binding:"required"`
	ConstructionSite string    `gorm:"type:varchar(1024)" json:"construction_site" binding:"required"`
	Quantity         uint      `gorm:"type:int unsigned" json:"quantity" binding:"required"`
	Price            uint      `gorm:"type:int unsigned" json:"price" binding:"required"`
	Charged          *bool     `gorm:"type:boolean" json:"charged" binding:"required"`
	Remark           string    `gorm:"type:varchar(4096)" json:"remark"`
}

func (WorkRecord) TableName() string {
	return "work_record"
}
