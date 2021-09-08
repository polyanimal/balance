package delivery

import (
	"Balance/internal/balance"
	"Balance/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlterFundsRequest struct {
	ID string `json:"id"`
	Funds string `json:"funds"`
}

type Handler struct {
	useCase balance.UseCase
}

func NewHandler(balanceUC balance.UseCase) *Handler {
	return &Handler{useCase: balanceUC}
}

func (h *Handler) AlterFunds(ctx *gin.Context) {
	req := new(AlterFundsRequest)
	err := ctx.BindJSON(req)

	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	currency := ctx.Param("currency")

	funds, err := strconv.Atoi(req.Funds)

	err = h.useCase.AlterFunds(req.ID, funds, currency)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) GetBalance(ctx *gin.Context) {

}

func (h *Handler) TransferFunds (ctx *gin.Context) {

}





