package delivery

import (
	"Balance/internal/balance"
	"Balance/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlterFundsRequest struct {
	Id    string `json:"id"`
	Funds string `json:"funds"`
}

type TransferRequest struct {
	IdFrom string `json:"id_from"`
	IdTo   string `json:"id_to"`
	Funds  string `json:"funds"`
}

type TransactionRequest struct {
	UserId  string `json:"user_id"`
	Sort    string `json:"sort"`
	Order   string `json:"order"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

type Handler struct {
	useCase balance.UseCase
}

func NewHandler(balanceUC balance.UseCase) *Handler {
	return &Handler{useCase: balanceUC}
}

func (h *Handler) AlterFunds(ctx *gin.Context) {
	req := new(AlterFundsRequest)

	if err := ctx.BindJSON(req); err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	funds, err := strconv.Atoi(req.Funds)
	if err != nil {
		msg := "Failed to cast funds to integer"
		util.RespondWithError(ctx, http.StatusBadRequest, msg)
		return
	}

	if err = h.useCase.AlterFunds(req.Id, funds); err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) GetBalance(ctx *gin.Context) {
	id := ctx.Param("id")
	currency := ctx.Query("currency")

	userBalance, err := h.useCase.GetBalance(id, currency)
	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, userBalance)
}

func (h *Handler) TransferFunds(ctx *gin.Context) {
	req := new(TransferRequest)

	if err := ctx.BindJSON(req); err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	funds, err := strconv.Atoi(req.Funds)
	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.useCase.TransferFunds(req.IdFrom, req.IdTo, funds); err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetTransactions(ctx *gin.Context) {
	req := new(TransactionRequest)
	if err := ctx.BindJSON(req); err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}


}
