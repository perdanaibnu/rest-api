package models

type Item struct {
	ItemId   int64  `gorm:"primaryKey" json:"id"`
	ItemCode string `gorm:"type:varchar(20)" json:"item_code"`
	ItemName string `gorm:"type:varchar(100)" json:"item_name"`
}
