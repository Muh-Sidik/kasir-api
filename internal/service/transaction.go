package service

import (
	"github.com/Muh-Sidik/kasir-api/internal/model"
	"github.com/Muh-Sidik/kasir-api/internal/model/dto"
	"github.com/Muh-Sidik/kasir-api/internal/repository"
)

type TransactionService interface {
	CreateCheckout(items []dto.CheckoutItem) (*model.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{
		repo: repo,
	}
}

func (t *transactionService) CreateCheckout(items []dto.CheckoutItem) (*model.Transaction, error) {
	return t.repo.CreateTransaction(items)
}
