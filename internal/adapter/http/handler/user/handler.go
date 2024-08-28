package userhandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sotoon/internal/adapter/http/request"
	"sotoon/internal/core/contract"
)

type Handler struct {
	userSvc contract.UserService
}

func New(userSvc contract.UserService) *Handler {
	return &Handler{
		userSvc: userSvc,
	}
}

func (r *Handler) CreateUser(ctx *gin.Context) {
	var req request.CreateUserRequest
	if ctx.BindJSON(&req) != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)

		return
	}

	user, err := r.userSvc.Create(ctx, req.MapToDTO())
	if err != nil {
		_ = ctx.Error(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusCreated, user)
}
