package dadata

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/jirawat-rackz/da-mock-api/pkg/model"
)

// IDAData is an interface for all data
type IDAData interface {
	// GetCoinList is a function to get coin list
	GetCoinList(paginate model.Pagination) ([]model.Coin, error)

	// GetCoinByID is a function to get coin by id
	GetCoinByID(ID string) (model.Coin, error)
}

// DAData is a struct for all data
type DAData struct {
	CoinData []model.Coin
}

// NewDAData is a function to create new DAData
func NewDAData() (IDAData, error) {
	var coins []model.Coin

	// Read JSON file
	file, err := os.ReadFile("mock-coins.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON
	if err := json.Unmarshal(file, &coins); err != nil {
		return nil, err
	}

	return &DAData{
		CoinData: coins,
	}, nil
}

func (d DAData) GetCoinList(paginate model.Pagination) ([]model.Coin, error) {

	firstIndex := (paginate.Page - 1) * paginate.Limit
	offset := paginate.Page * paginate.Limit

	if firstIndex > len(d.CoinData) {
		return []model.Coin{}, errors.New("no data")
	}

	if offset > len(d.CoinData) {
		offset = len(d.CoinData)
	}

	return d.CoinData[firstIndex:offset], nil
}

func (d DAData) GetCoinByID(ID string) (model.Coin, error) {
	for _, coin := range d.CoinData {
		if coin.ID == ID {
			return coin, nil
		}
	}

	return model.Coin{}, nil
}
