package repository

import (
	"github.com/IrvanWijayaSardam/Bengkel/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(b *entity.Transaction) entity.Transaction
	All(idUser string) []entity.Transaction
	DeleteTransaction(b *entity.Transaction)
	FindTransactionById(UserID uint64) entity.Transaction
}

type transactionConnection struct {
	connection *gorm.DB
}

// FindTransactionById implements TransactionRepository
func (db *transactionConnection) FindTransactionById(UserID uint64) entity.Transaction {
	var transaction entity.Transaction
	db.connection.Preload("User").Find(&transaction, UserID)
	return transaction
}

// All implements TransactionRepository
func (db *transactionConnection) All(idUser string) []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Preload("User").Where("user_id = ?", idUser).Find(&transactions)
	return transactions
}

// DeleteTransaction implements TransactionRepository
func (db *transactionConnection) DeleteTransaction(b *entity.Transaction) {
	db.connection.Delete(&b)
}

// InsertTransaction implements TransactionRepository
func (db *transactionConnection) InsertTransaction(b *entity.Transaction) entity.Transaction {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return *b
}

func NewTransactionRepository(dbConn *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: dbConn,
	}
}
