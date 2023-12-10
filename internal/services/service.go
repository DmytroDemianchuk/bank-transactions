package services

import (
	"context"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
)

type IRepoBank interface {
	GetFilteredData(ctx context.Context, input domain.FilterSearchInput) ([]domain.Transaction, error)
	InsertTransactions(ctx context.Context, transactions []domain.Transaction) error
}
type IRepoRemote interface {
	Get(ctx context.Context, from, to *int) ([]domain.Transaction, error)
}

type Services struct {
	ServicesBank   *ServiceBank
	ServicesRemote *ServiceRemoteCSV
}

func New(repoBank IRepoBank, repoRemote IRepoRemote) *Services {
	return &Services{
		ServicesBank:   NewBankServices(repoBank),
		ServicesRemote: NewRemoteServices(repoRemote),
	}
}
