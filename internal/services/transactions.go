package services

import (
	"context"
	"sync"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
)

type ServiceBank struct {
	repo           IRepoBank
	downloadStatus sync.Map
}

func NewBankServices(repo IRepoBank) *ServiceBank {
	se := ServiceBank{repo: repo}
	se.downloadStatus.Store("downloadStatus", domain.Unknown)
	return &se
}

func (s *ServiceBank) GetFilteredData(ctx context.Context, input domain.FilterSearchInput) ([]domain.Transaction, error) {
	return s.repo.GetFilteredData(ctx, input)
}

func (s *ServiceBank) FetchExternTransactions(ctx context.Context, url string) (domain.Status, error) {
	downloadStatus := domain.Unknown
	if ds, ok := s.downloadStatus.Load("downloadStatus"); ok {
		downloadStatus = ds.(domain.Status)
	}

	if url == "" {
		return downloadStatus, nil
	}

	if downloadStatus == domain.Processing || downloadStatus == domain.Skip {
		if downloadStatus == domain.Processing {
			s.downloadStatus.Store("downloadStatus", domain.Skip)
			return domain.Skip, nil
		}
		return downloadStatus, nil
	}

	// Request a list of transactions from an external service via REST
	s.downloadStatus.Store("downloadStatus", domain.Processing)
	go s.workerPoolDownloadTransactions(ctx, url)

	return domain.Processing, nil
}
