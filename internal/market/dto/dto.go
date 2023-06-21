package dto

type TradeInput struct {
	OrderID       string  `json:"order_id"`
	InvesorID     string  `json:"invesor_id"`
	AssetID       string  `json:"asset_id"`
	CurrentShares int     `json:"current_Shares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"OrderType"`
}

type OrderOutput struct {
	OrderID           string               `json:"order_id"`
	InvesorID         string               `json:"invesor_id"`
	AssetID           string               `json:"asset_id"`
	OrderType         string               `json:"OrderType"`
	Status            string               `json:"status"`
	Partial           int                  `json:"partial"`
	Shares            int                  `json:"shares"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerID       string  `json:"buyer_id"`
	SellerID      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}
