package model

type Customer struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerName string `gorm:"type:varchar(256)" json:"customer_name"`
}

func (Customer) TableName() string {
	return "customer_name"
}
