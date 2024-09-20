package price

import "go-otlp/repositories/price"

type priceService struct {
	priceRepo price.PriceRepository
}

func NewPriceService(priceRepo price.PriceRepository) PriceService {
	return priceService{priceRepo: priceRepo}
}

func (s priceService) GetPrices() ([]PriceResponse, error) {
	prices, err := s.priceRepo.GetPrices()

	if err != nil {
		return nil, err
	}

	priceResponses := []PriceResponse{}

	for _, price := range prices {
		priceResponse := PriceResponse{
			Id:    price.Id,
			Value: price.Value,
		}
		priceResponses = append(priceResponses, priceResponse)
	}

	return priceResponses, nil
}

func (s priceService) GetPrice(id int) (*PriceResponse, error) {
	price, err := s.priceRepo.GetPrice(id)
	if err != nil {
		return nil, err
	}

	priceReponse := PriceResponse{
		Id:    price.Id,
		Value: price.Value,
	}

	return &priceReponse, nil
}

func (s priceService) AddPrice(request PriceRequest) (int, error) {
	price := price.Price{
		Value: request.Value,
	}

	result, err := s.priceRepo.AddPrice(price)

	if err != nil {
		return 0, err
	}

	return result, nil

}
