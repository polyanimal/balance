package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/balance/internal/balance"
	"github.com/polyanimal/balance/internal/models"
	"github.com/polyanimal/balance/util"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase balance.UseCase
}

func NewHandler(balanceUC balance.UseCase) *Handler {
	return &Handler{useCase: balanceUC}
}

func (h *Handler) AlterFunds(ctx *gin.Context) {
	req := new(models.AlterFundsRequest)

	if err := ctx.BindJSON(req); err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	funds, err := strconv.Atoi(req.Funds)
	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, "Failed to cast funds to integer")
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

	w := userBalance / 100
	c := userBalance - w * 100
	v := strconv.Itoa(w) + "." + strconv.Itoa(c)

	if currency == "" {
		currency = "RUB"
	}

	resp := models.BalanceResponse{
		Value:    v,
		Currency: currency,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) TransferFunds(ctx *gin.Context) {
	req := new(models.TransferRequest)

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
	req := new(models.TransactionsRequest)
	if err := ctx.BindJSON(req); err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if req.Page <= 0 {
		util.RespondWithError(ctx, http.StatusBadRequest, "invalid page")
		return
	}

	transactions, err := h.useCase.GetTransactions(*req)
	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
