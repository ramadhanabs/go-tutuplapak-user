package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"go-tutuplapak-user/models"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	EmailExists(email string) (bool, error)
	PhoneExists(phone string) (bool, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.Password,
		&user.FileID, &user.FileURI, &user.FileThumbnailURI,
		&user.BankAccountName, &user.BankAccountHolder, &user.BankAccountNumber,
		&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error querying user: %w", err)
	}

	return &user, nil
}

func (r *userRepository) FindByPhone(phone string) (*models.User, error) {
	query := "SELECT * FROM users WHERE phone = $1"
	row := r.db.QueryRow(query, phone)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.Password,
		&user.FileID, &user.FileURI, &user.FileThumbnailURI,
		&user.BankAccountName, &user.BankAccountHolder, &user.BankAccountNumber,
		&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error querying user: %w", err)
	}

	return &user, nil
}

func (r *userRepository) EmailExists(email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	var exists bool
	if err := r.db.QueryRow(query, email).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

func (r *userRepository) PhoneExists(phone string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE phone = $1)"
	var exists bool
	if err := r.db.QueryRow(query, phone).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (email, phone, password, bank_account_name, bank_account_holder, bank_account_number) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, user.Email, user.Phone, user.Password, user.BankAccountName, user.BankAccountHolder, user.BankAccountNumber)
	return err
}
