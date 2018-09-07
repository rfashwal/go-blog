package model

type Account struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ServedBy string `json:"served_by"`
	Quote    Quote  `json:"quote"`
}

type Quote struct {
	Text     string `json:"quote"`
	ServedBy string `json:"served_by"`
	Language string `json:"language"`
}
