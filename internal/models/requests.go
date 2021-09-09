package models

type AlterFundsRequest struct {
	Id    string `json:"id"`
	Funds string `json:"funds"`
}

type TransferRequest struct {
	IdFrom string `json:"id_from"`
	IdTo   string `json:"id_to"`
	Funds  string `json:"funds"`
}

type TransactionsRequest struct {
	UserId  string `json:"user_id"`
	Sort    string `json:"sort"`
	Order   string `json:"order"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

