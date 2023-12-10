package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
	"github.com/dmytrodemianchuk/bank-transactions/pkg/restclient"
	"github.com/gocarina/gocsv"
	"github.com/sirupsen/logrus"
)

const transactionsCount = 2
const workerCount = 3

var wg *sync.WaitGroup

type job struct {
	mu           sync.Mutex
	From         int
	To           int
	Transactions int
	Status       domain.Status
}

func (j *job) getFromTo() (int, int) {
	j.mu.Lock()
	defer j.mu.Unlock()

	if j.From == 0 && j.To == 0 {
		j.From = 1
		j.To = transactionsCount
	} else {
		j.From = j.From + transactionsCount
		j.To = j.To + transactionsCount
	}
	return j.From, j.To
}

func (j *job) getSetStatus(tr int, status domain.Status) {
	j.mu.Lock()
	j.Status = status
	j.Transactions = j.Transactions + tr
	j.mu.Unlock()
}

func (s *ServiceBank) workerGetTransactionsRemoteURL(ctx context.Context, url string) (int, error) {
	restClient, err := restclient.NewClient(time.Second * 180)
	if err != nil {
		logrus.Error(err)
		return 0, errors.New("error create http client")
	}

	// REST request
	in_csv, err := restClient.Get(url)
	if err != nil {
		logrus.Error(err)
		return 0, errors.New("error REST request")
	}

	if string(in_csv) == `{"error":"no data"}` {
		return 0, nil
	}

	var transactions []domain.Transaction
	// UnmarshalBytes parses the CSV from the bytes in the interface.
	err = gocsv.UnmarshalBytes(in_csv, &transactions)
	if err != nil {
		logrus.Error(err)
		return 0, errors.New("error unmarshal csv transactions")
	}

	if len(transactions) == 0 {
		return 0, nil
	}

	err = s.repo.InsertTransactions(ctx, transactions)
	if err != nil {
		logrus.Error(err)
		return 0, errors.New("error db insert transactions")
	}

	return len(transactions), nil
}

func (s *ServiceBank) worker(wid int, ctx context.Context, url string, job *job) {
	for {
		from, to := job.getFromTo()
		fmt.Printf("Worker #%d start: Get transactions from %d  to %d\n", wid, from, to)

		{
			{
				{
					{
						{
							{
								time.Sleep(time.Second * 2) // ----------------------- Timeout for fake -----------------------
							}
						}
					}
				}
			}
		}

		url_from_to := url + "?from=" + strconv.Itoa(from) + "&to=" + strconv.Itoa(to)
		trns, err := s.workerGetTransactionsRemoteURL(ctx, url_from_to)

		if err != nil {
			fmt.Printf("Worker #%d done: DownloadError\n", wid)
			job.getSetStatus(0, domain.DownloadError)
			break
		}
		if trns == 0 {
			fmt.Printf("Worker #%d done: DownloadOk\n", wid)
			job.getSetStatus(0, domain.DownloadOk)
			break
		}

		job.getSetStatus(trns, domain.Processing)
	}

	wg.Done()
}

func (s *ServiceBank) workerPoolDownloadTransactions(ctx context.Context, url string) {
	fmt.Printf("\n\n------------------------------ Worker Pool Start!\n")

	startTime := time.Now()

	var job job

	wg = &sync.WaitGroup{}
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go s.worker(i, ctx, url, &job)
	}
	wg.Wait()

	s.downloadStatus.Store("downloadStatus", job.Status)

	fmt.Printf("------------------------------ Worker Pool DONE: %s\n", string(job.Status)+": "+strconv.Itoa(job.Transactions)+" transactions loaded")
	fmt.Printf("Download time elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}
