package models

//type CurrencyConversionResponse struct {
//	Success bool `json:"success"`
//	Historical bool `json:"historical"`
//	Date string `json:"date"`
//	Timestamp int64 `json:"timestamp"`
//}

type BalanceResponse struct {
	Value string
	Currency string
}