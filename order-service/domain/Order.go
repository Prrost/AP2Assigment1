package domain

type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	ProductID int    `json:"productId"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
}
