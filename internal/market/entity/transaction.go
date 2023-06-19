package entity

import "time"

type Transaction struct {
	ID string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares int
	Price float64
	Total float64
	Date time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) * Transaction {
	total := float64(shares) * price
	
}