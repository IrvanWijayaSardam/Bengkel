package dto

type TransactionCreateDTO struct {
	UserID uint64                       `json:"userid" form:"userid" binding:"required"`
	Date   string                       `json:"date" form:"date" binding:"required"`
	Detail []TransactionDetailCreateDTO `json:"detail" form:"detail" binding:"required,dive"`
}

type TransactionDetailCreateDTO struct {
	Description string `json:"description" form:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
}
