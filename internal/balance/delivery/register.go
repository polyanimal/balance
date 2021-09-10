package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/polyanimal/balance/internal/balance"
)

func RegisterHTTPEndpoints(router *gin.Engine, balanceUC balance.UseCase) {
	handler := NewHandler(balanceUC)

	router.GET("/balance/:id", handler.GetBalance)
	router.POST("/balance", handler.AlterFunds)
	router.POST("/transfer", handler.TransferFunds)
	router.GET("/transactions", handler.GetTransactions)
}
