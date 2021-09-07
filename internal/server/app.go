package server

import (
	"Balance/internal/balance"
	"net/http"
)

type App struct {
	server    *http.Server
	balanceUC balance.UseCase
}
