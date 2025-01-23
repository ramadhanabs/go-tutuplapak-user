package utils

import (
	"net/http"

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
