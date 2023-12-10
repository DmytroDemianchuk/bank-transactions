package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
)

const tableSourceCSV = "sourceCSV"

type RepoRemote struct {
	db *sqlx.DB
}

func NewRemoteRepo(db *sqlx.DB) *RepoRemote {
	return &RepoRemote{db: db}
}

func (r *RepoRemote) Get(ctx context.Context, from, to *int) ([]domain.Transaction, error) {
	var sourceCSV []domain.Transaction

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if from != nil {
		setValues = append(setValues, fmt.Sprintf("transaction_id>=$%d", argId))
		args = append(args, *from)
		fmt.Printf("%v", *from)
		argId++
	}

	if to != nil {
		setValues = append(setValues, fmt.Sprintf("transaction_id<=$%d", argId))
		args = append(args, *to)
		fmt.Printf("%v", *to)
		// argId++
	}
	setQuery := strings.Join(setValues, " AND ")

	query := fmt.Sprintf(`SELECT * FROM %s`, tableSourceCSV)
	if setQuery != "" {
		query = query + " WHERE " + setQuery
	}

	if err := r.db.SelectContext(ctx, &sourceCSV, query, args...); err != nil {
		return nil, err
	}

	return sourceCSV, nil
}
