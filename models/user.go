package models

import "database/sql"

type User struct {
	ID                int            `json:"id"`
	Email             sql.NullString `json:"email"`
	Phone             sql.NullString `json:"phone"`
	Password          string         `json:"password"`
	FileID            string         `json:"file_id"`
	FileURI           string         `json:"file_uri"`
	FileThumbnailURI  string         `json:"file_thumbnail_uri"`
	BankAccountName   string         `json:"bank_account_name"`
	BankAccountHolder string         `json:"bank_account_holder"`
	BankAccountNumber string         `json:"bank_account_number"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
}
