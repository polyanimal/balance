package delivery

import (
	"Balance/internal/balance"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase balance.UseCase
}

func NewHandler(balanceUC balance.UseCase) *Handler {
	return &Handler{useCase: balanceUC}
}

func (h *Handler) AlterFunds(ctx *gin.Context) {

}

func (h *Handler) GetBalance(ctx *gin.Context) {

}

func (h *Handler) TransferFunds (ctx *gin.Context) {

}





