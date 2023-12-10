package domain

import "time"

type UrlInput struct {
	URL *string `json:"url,omitempty" example:"http://localhost:8080/api/v1/get_csv_mock_remote_service"`
}

// пошук по
type FilterSearchInput struct {
	TransactionId    *int    `json:"transaction_id,omitempty"    example:"18"`
	TerminalId       []int   `json:"terminal_id,omitempty"       example:"3521,3522,3523,3524,3525,3526,3527,3528,3529"`
	Status           *string `json:"status,omitempty"            validate:"omitempty,oneof=accepted declined" swaggertype:"string" enums:"accepted,declined" example:"accepted"`
	PaymentType      *string `json:"payment_type,omitempty"      validate:"omitempty,oneof=cash card" swaggertype:"string" enums:"cash,card" example:"cash"`
	Period           *Period `json:"period,omitempty"            validate:"omitempty"`
	PaymentNarrative *string `json:"payment_narrative,omitempty" example:"contract for the provision of services A11/27123"`
}
type Period struct {
	From *time.Time `json:"from" validate:"required"  example:"2023-08-23T11:56:00.000Z"` // example: from 2023-08-12, to 2023-09-01 must return all transactions for the specified period
	To   *time.Time `json:"to"   validate:"required"  example:"2023-08-24T00:00:00.000Z"`
}

type Transaction struct {
	ID                 int       `json:"-" csv:"-"           db:"id"`
	TransactionId      int       `json:"transaction_id"      db:"transaction_id"`
	RequestId          int       `json:"request_id"          db:"request_id"`
	TerminalId         int       `json:"terminal_id"         db:"terminal_id"`
	PartnerObjectId    int       `json:"partner_object_id"   db:"partner_object_id"`
	AmountTotal        float32   `json:"amount_total"        db:"amount_total"         example:"1.23"`
	AmountOriginal     float32   `json:"amount_original"     db:"amount_original"      example:"1.23"`
	CommissionPS       float32   `json:"commission_ps"       db:"commission_ps"        example:"1.23"`
	CommissionClient   float32   `json:"commission_client"   db:"commission_client"    example:"1.23"`
	CommissionProvider float32   `json:"commission_provider" db:"commission_provider"  example:"1.23"`
	DateInput          time.Time `json:"date_input"          db:"date_input"`
	DatePost           time.Time `json:"date_post"           db:"date_post"`
	Status             string    `json:"status"              db:"status"`
	PaymentType        string    `json:"payment_type"        db:"payment_type"`
	PaymentNumber      string    `json:"payment_number"      db:"payment_number"`
	ServiceId          int       `json:"service_id"          db:"service_id"`
	Service            string    `json:"service"             db:"service"`
	PayeeId            int       `json:"payee_id"            db:"payee_id"`
	PayeeName          string    `json:"payee_name"          db:"payee_name"`
	PayeeBankMfo       int       `json:"payee_bnank_mfo"     db:"payee_bnank_mfo"`
	PayeeBankAccount   string    `json:"payee_bnank_account" db:"payee_bnank_account"`
	PaymentNarrative   string    `json:"payment_narrative"   db:"payment_narrative"`
}
