package price

type Price struct {
	Id    int `gorm:"primary"`
	Value float64
}

func (p *Price) TableName() string {
	return "pricing.price"
}

type PriceRepository interface {
	GetPrices() ([]Price, error)
	GetPrice(int) (*Price, error)
	AddPrice(Price) (int, error)
}
