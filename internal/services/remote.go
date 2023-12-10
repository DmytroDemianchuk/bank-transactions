package services

import (
	"context"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
)

type ServiceRemoteCSV struct {
	repo IRepoRemote
}

func NewRemoteServices(repo IRepoRemote) *ServiceRemoteCSV {
	return &ServiceRemoteCSV{repo: repo}
}

func (s *ServiceRemoteCSV) Get(ctx context.Context, from, to *int) ([]domain.Transaction, error) {
	return s.repo.Get(ctx, from, to)
}
