package controllers

import (
	"go-tutuplapak-user/services"
	"go-tutuplapak-user/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

type LoginPhoneResponse struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginEmailResponse struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) LoginWithEmail(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=32"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondValidationError(ctx, err)
		return
	}

	user, token, err := c.authService.LoginWithEmail(req.Email, req.Password)
	if err != nil {
		utils.RespondError(ctx, http.StatusNotFound, err.Error())
		return
	}

	userResponse := utils.ToUserResponse(user)

	utils.RespondJSON(ctx, http.StatusOK, LoginEmailResponse{
		Email: userResponse.Email,
		Phone: userResponse.Phone,
		Token: token,
	})
}

func (c *AuthController) LoginWithPhone(ctx *gin.Context) {
	var req struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=32"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondValidationError(ctx, err)
		return
	}

	user, token, err := c.authService.LoginWithPhone(req.Phone, req.Password)
	if err != nil {
		utils.RespondError(ctx, http.StatusNotFound, err.Error())
		return
	}

	userResponse := utils.ToUserResponse(user)

	utils.RespondJSON(ctx, http.StatusOK, LoginPhoneResponse{
		Phone: userResponse.Phone,
		Email: userResponse.Email,
		Token: token,
	})
}

func (c *AuthController) RegisterWithEmail(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=32"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondValidationError(ctx, err)
		return
	}

	user, token, err := c.authService.RegisterWithEmail(req.Email, req.Password)
	if err != nil {
		status := http.StatusConflict
		if err.Error() == "email already exists" {
			status = http.StatusConflict
		}
		utils.RespondError(ctx, status, err.Error())
		return
	}

	userResponse := utils.ToUserResponse(user)

	utils.RespondJSON(ctx, http.StatusCreated, gin.H{"email": userResponse.Email, "token": token})
}

func (c *AuthController) RegisterWithPhone(ctx *gin.Context) {
	var req struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=32"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondValidationError(ctx, err)
		return
	}

	user, token, err := c.authService.RegisterWithPhone(req.Phone, req.Password)
	if err != nil {
		status := http.StatusConflict
		if err.Error() == "phone already exists" {
			status = http.StatusConflict
		} else {
			status = http.StatusBadRequest
		}
		utils.RespondError(ctx, status, err.Error())
		return
	}

	userResponse := utils.ToUserResponse(user)

	utils.RespondJSON(ctx, http.StatusCreated, gin.H{"phone": userResponse.Phone, "token": token})
}
