package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondJSON sends a JSON response with the specified status code and payload.
func RespondJSON(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}

// RespondError sends an error response with a specified status code and message.
func RespondError(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{"error": message})
}

// RespondValidationError handles validation errors.
func RespondValidationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
