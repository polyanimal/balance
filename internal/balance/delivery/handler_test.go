package delivery

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/polyanimal/balance/internal/balance/mocks"
	"github.com/polyanimal/balance/internal/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMoviesHandlers(t *testing.T) {
	r := gin.Default()
	ctrl := gomock.NewController(t)
	uc := mocks.NewMockUseCase(ctrl)

	RegisterHTTPEndpoints(r, uc)
	a := models.AlterFundsRequest{
		Id: "test",
		Funds: "10",
	}

	body, err := json.Marshal(a)
	assert.NoError(t, err)

	t.Run("AlterFunds", func(t *testing.T) {
		uc.EXPECT().AlterFunds("test", 10).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/balance", bytes.NewBuffer(body))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GetBalance", func(t *testing.T) {
		uc.EXPECT().GetBalance("test", "").Return(1000, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/balance/test", bytes.NewBuffer(body))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
