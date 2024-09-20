package price

type PriceResponse struct {
	Id    int     `json:"id"`
	Value float64 `json:"value"`
}

type PriceRequest struct {
	Value float64 `json:"value"`
}

type PriceService interface {
	GetPrices() ([]PriceResponse, error)
	GetPrice(int) (*PriceResponse, error)
	AddPrice(PriceRequest) (int, error)
}
