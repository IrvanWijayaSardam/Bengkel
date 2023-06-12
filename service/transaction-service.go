package service

import (
	"fmt"
	"log"

	"github.com/IrvanWijayaSardam/Bengkel/dto"
	"github.com/IrvanWijayaSardam/Bengkel/entity"
	"github.com/IrvanWijayaSardam/Bengkel/repository"

	"github.com/mashingan/smapping"
)

type TransactionService interface {
	InsertTransaction(b dto.TransactionCreateDTO) entity.Transaction
	Delete(b entity.Transaction)
	IsAllowedToEdit(userID string, transactionID uint64) bool
	All(idUser string) []entity.Transaction
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

// All implements TransactionService
func (service *transactionService) All(idUser string) []entity.Transaction {
	return service.transactionRepository.All(idUser)
}

// Delete implements TransactionService
func (service *transactionService) Delete(b entity.Transaction) {
	service.transactionRepository.DeleteTransaction(&b)
}

// InsertTransaction implements TransactionService
func (service *transactionService) InsertTransaction(b dto.TransactionCreateDTO) entity.Transaction {
	trx := entity.Transaction{}
	err := smapping.FillStruct(&trx, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.transactionRepository.InsertTransaction(&trx)
	return res
}

// IsAllowedToEdit implements TransactionService
func (service *transactionService) IsAllowedToEdit(userID string, transactionID uint64) bool {
	b := service.transactionRepository.FindTransactionById(transactionID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepo,
	}
}
