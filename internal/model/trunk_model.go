package model

type TrunkModel struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	TrunkModel string `gorm:"type:varchar(256)" json:"trunk_model"`
}

func (TrunkModel) TableName() string {
	return "trunk_model"
}
