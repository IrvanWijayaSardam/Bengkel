package entity

type Transaction struct {
	ID     uint64              `gorm:"primary_key:auto_increment" json:"id"`
	UserID uint64              `gorm:"index" json:"user_id"`
	Date   string              `json:"date"`
	Detail []TransactionDetail `gorm:"foreignkey:TransactionID" json:"detail"`
}

type TransactionDetail struct {
	ID            uint64 `gorm:"primary_key:auto_increment" json:"id"`
	TransactionID uint64 `gorm:"index" json:"transaction_id"`
	Description   string `gorm:"type:varchar(255)" json:"description"`
	Price         uint64 `gorm:"type:varchar(255)" json:"price"`
}
