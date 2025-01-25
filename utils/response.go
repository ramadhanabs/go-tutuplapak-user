package utils

import (
	"database/sql"
	"go-tutuplapak-user/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func RespondJSON(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}

func RespondError(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{"error": message})
}

func RespondValidationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func IsValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+(\d{1,15})$`) // Matches international phone number with '+' prefix
	return re.MatchString(phone)
}

type userResponse struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Password          string `json:"password"`
	FileID            string `json:"file_id"`
	FileURI           string `json:"file_uri"`
	FileThumbnailURI  string `json:"file_thumbnail_uri"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountHolder string `json:"bank_account_holder"`
	BankAccountNumber string `json:"bank_account_number"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func ToUserResponse(user *models.User) *userResponse {
	return &userResponse{
		ID:                user.ID,
		Email:             nullableToString(user.Email),
		Phone:             nullableToString(user.Phone),
		Password:          user.Password,
		FileID:            user.FileID,
		FileURI:           user.FileURI,
		FileThumbnailURI:  user.FileThumbnailURI,
		BankAccountName:   user.BankAccountName,
		BankAccountHolder: user.BankAccountHolder,
		BankAccountNumber: user.BankAccountNumber,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
}

func nullableToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
