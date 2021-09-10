package usecase

import (
	"encoding/json"
	"errors"
	"github.com/polyanimal/balance/internal/balance"
	"github.com/polyanimal/balance/internal/models"
	"io/ioutil"
	"net/http"
	"time"
)

type BalanceUseCase struct {
	repo balance.Repository
}

func NewBalanceUC(repo balance.Repository) *BalanceUseCase {
	return &BalanceUseCase{
		repo: repo,
	}
}

func (uc *BalanceUseCase) GetBalance(id string, currency string) (int, error) {
	funds, err := uc.repo.GetBalance(id)
	if err != nil {
		return 0, err
	}

	if currency != "" && currency != "RUB" {
		url := "https://www.cbr-xml-daily.ru/daily_json.js"

		response, err := http.Get(url)
		if err != nil {
			return 0, errors.New("failed to convert funds")
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return 0, errors.New("failed to read conversion data")
		}

		var v interface{}
		if err := json.Unmarshal(responseData, &v); err != nil {
			return 0, errors.New("failed to unmarshal conversion response")
		}

		data := v.(map[string]interface{})
		rate := 1.0 / data["Valute"].(map[string]interface{})[currency].(map[string]interface{})["Value"].(float64)

		funds = int(float64(funds) * rate)

		//
		//if data["success"].(bool) {
		//	balance = data["result"].(int)
		//} else {
		//	return 0, errors.New("funds conversion unsuccessful")
		//}
	}

	return funds, nil
}

func (uc *BalanceUseCase) AlterFunds(id string, funds int) error {
	t := time.Now()

	if err := uc.repo.AlterFunds(id, funds); err != nil {
		return err
	}

	var operation string
	if funds >= 0 {
		operation = "funds added"
	} else {
		operation = "funds taken"
	}

	if err := uc.repo.RecordTransaction(operation, id, "", funds, t); err != nil {
		return errors.New("error while recording a transaction")
	}

	return nil
}

func (uc *BalanceUseCase) TransferFunds(idFrom, idTo string, funds int) error {
	t := time.Now()

	//Предполагается что пользователь может только перевести свои средства другому пользователю, а не наоборот
	if funds < 0 {
		return errors.New("invalid operation")
	}

	if err := uc.repo.TransferFunds(idFrom, idTo, funds); err != nil {
		return err
	}

	if err := uc.repo.RecordTransaction("funds transfer", idFrom, idTo, funds, t); err != nil {
		return errors.New("error while recording a transaction")
	}

	return nil
}

func (uc *BalanceUseCase) GetTransactions(request models.TransactionsRequest) ([]models.Transaction, error) {
	return uc.repo.GetTransactions(request.UserId, request.Order, request.Sort, request.Page, request.PerPage)
}
