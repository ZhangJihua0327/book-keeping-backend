package model

import "time"

type WorkRecord struct {
	ID               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	RecordID         string    `gorm:"type:varchar(256)" json:"record_id"`
	TrunkModel       string    `gorm:"type:varchar(256)" json:"trunk_model"`
	Date             time.Time `gorm:"type:date" json:"date"` // Using time.Time for date, might need custom scanner or just "2006-01-02" string if purely JSON handling preference, but time.Time is standard for GORM
	CustomerName     string    `gorm:"type:varchar(256)" json:"customer_name"`
	ConstructionSite string    `gorm:"type:varchar(1024)" json:"construction_site"`
	Quantity         uint      `gorm:"type:int unsigned" json:"quantity"`
	Price            uint      `gorm:"type:int unsigned" json:"price"`
	Charged          bool      `gorm:"type:boolean" json:"charged"`
	Remark           string    `gorm:"type:varchar(4096)" json:"remark"`
}

func (WorkRecord) TableName() string {
	return "work_record"
}
