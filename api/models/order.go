package models

type Order struct {
	PizzaID  int `json:"pizza_id"`
	Quantity int `json:"quantity"`
	Total    int `json:"total"`
}
