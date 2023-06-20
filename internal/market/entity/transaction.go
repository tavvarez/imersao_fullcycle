package entity

import (
	"time"
	"github.com/google/uuid"
)

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
	return &Transaction{
		ID: uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares: shares,
		Price: price,
		Total: total,
		Date: time.Now(),
	}
}

func (t *Transaction) CalculateTotal(shares int, price float64) {
	t.total = float64(t.shares) * t.price
}

func (t *Transaction) CloseBuyOrderTransaction() {
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) CloseSellOrderTransaction() {
	if t.SellingOrder.PendingShares == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) AddBuyOrderPendingShares(shares int) {
	t.BuyingOrder.PendingShares += shares
}

func (t *Transaction) AddSellOrderPendingShares(shares int) {
	t.SellingOrder.PendingShares += shares
}