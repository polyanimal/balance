package delivery

import (
	"Balance/internal/balance"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, balanceUC balance.UseCase) {
	handler := NewHandler(balanceUC)

	router.GET("/balance", handler.GetBalance)
	router.POST("/payment", handler.AlterFunds)
	router.POST("/transfer", handler.TransferFunds)

}
