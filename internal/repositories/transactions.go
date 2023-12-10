package repositories

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/dmytrodemianchuk/bank-transactions/internal/domain"
)

const tableTransactions = "transactions"

type RepoBank struct {
	db *sqlx.DB
}

func NewRepoBank(db *sqlx.DB) *RepoBank {
	return &RepoBank{db: db}
}

func (r *RepoBank) GetFilteredData(ctx context.Context, input domain.FilterSearchInput) ([]domain.Transaction, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.TransactionId != nil {
		setValues = append(setValues, fmt.Sprintf("transaction_id=$%d", argId))
		args = append(args, *input.TransactionId)
		argId++
	}

	if input.TerminalId != nil {
		var str []string
		for _, v := range input.TerminalId {
			args = append(args, v)
			str = append(str, "$"+strconv.Itoa(argId))
			argId++
		}
		setValues = append(setValues, fmt.Sprintf("terminal_id IN (%s)", strings.Join(str, ",")))
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	if input.PaymentType != nil {
		setValues = append(setValues, fmt.Sprintf("payment_type=$%d", argId))
		args = append(args, *input.PaymentType)
		argId++
	}

	if input.Period != nil {
		if input.Period.From != nil {
			setValues = append(setValues, fmt.Sprintf("date_post>=$%d", argId))
			args = append(args, *input.Period.From)
			fmt.Printf("%v", *input.Period.From)
			argId++
		}

		if input.Period.To != nil {
			setValues = append(setValues, fmt.Sprintf("date_post<=$%d", argId))
			args = append(args, *input.Period.To)
			fmt.Printf("%v", *input.Period.To)
			argId++
		}
	}

	if input.PaymentNarrative != nil {
		setValues = append(setValues, fmt.Sprintf("payment_narrative ILIKE $%d", argId))
		a := fmt.Sprintf("%%%s%%", *input.PaymentNarrative)
		args = append(args, a)
		// argId++
	}

	setQuery := strings.Join(setValues, " AND ")

	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s`, tableTransactions, setQuery)
	var csv []domain.Transaction
	if err := r.db.SelectContext(ctx, &csv, query, args...); err != nil {
		return nil, err
	}

	return csv, nil
}

func (r *RepoBank) InsertTransactions(ctx context.Context, transactions []domain.Transaction) error {
	query := fmt.Sprintf(`INSERT INTO %s 
		(transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bnank_mfo, payee_bnank_account, payment_narrative) 
		VALUES (:transaction_id, :request_id, :terminal_id, :partner_object_id, :amount_total, :amount_original, :commission_ps, :commission_client, :commission_provider, :date_input, :date_post, :status, :payment_type, :payment_number, :service_id, :service, :payee_id, :payee_name, :payee_bnank_mfo, :payee_bnank_account, :payment_narrative)`,
		tableTransactions)

	_, err := r.db.NamedExec(query, transactions)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
