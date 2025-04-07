package domain

type Object struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int64  `json:"amount"`
	Available bool   `json:"available"`
}
