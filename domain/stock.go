package domain

type Stock struct {
	Type  string
	Size  string
	Count int
	Price float64
}

type StockReq struct {
	Type  string
	Size  string
	Count int
	Price float64
}

type StockResp []StockReq

type StockService interface {
	CreateStock(req Stock) (Stock, error)
}

type StockRepository interface {
	CreateStock(req Stock) (Stock, error)
}
