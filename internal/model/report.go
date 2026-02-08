package model

type RevenueReport struct {
	TotalRevenue int64 `json:"total_revenue"`
}

type TransactionReport struct {
	TotalTransaction int64 `json:"total_transaksi"`
}

type TopProduct struct {
	ProductName  string `json:"nama"`
	QuantitySold int64  `json:"qty_terjual"`
}

type TopProductReport struct {
	RevenueReport
	TransactionReport
	Products TopProduct `json:"produk_terlaris"`
}
