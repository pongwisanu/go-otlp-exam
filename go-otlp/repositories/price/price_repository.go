package price

import "gorm.io/gorm"

type priceRepositoryDb struct {
	db *gorm.DB
}

func NewPriceRepositoryDb(db *gorm.DB) PriceRepository {
	return priceRepositoryDb{db: db}
}

func (r priceRepositoryDb) GetPrices() ([]Price, error) {
	prices := []Price{}
	result := r.db.Find(&prices)

	if result.Error != nil {
		return nil, result.Error
	}

	return prices, nil
}

func (r priceRepositoryDb) GetPrice(id int) (*Price, error) {
	price := Price{}
	result := r.db.First(&price, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &price, nil
}

func (r priceRepositoryDb) AddPrice(newPrice Price) (int, error) {

	result := r.db.Create(&newPrice)

	if result.Error != nil {
		return 0, result.Error
	}

	return newPrice.Id, nil
}
