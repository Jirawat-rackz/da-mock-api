package darepository

import (
	"github.com/jirawat-rackz/da-mock-api/data/dadata"
	"github.com/jirawat-rackz/da-mock-api/pkg/model"
)

// IDARepository is an interface for all repository
type IDARepository interface {
	// GetCoinByID is a function to get coin by id
	GetCoinByID(ID string) (model.Coin, error)

	// GetCoinList is a function to get coin list
	GetCoinList(pagination model.Pagination) ([]model.Coin, error)
}

// DARepository is a struct for all repository
type DARepository struct {
	DAData dadata.IDAData
}

// NewDARepository is a function to create new DARepository
func NewDARepository(daData dadata.IDAData) IDARepository {
	return &DARepository{
		DAData: daData,
	}
}

func (d DARepository) GetCoinList(pagination model.Pagination) ([]model.Coin, error) {
	return d.DAData.GetCoinList(pagination)
}

func (d DARepository) GetCoinByID(ID string) (model.Coin, error) {
	return d.DAData.GetCoinByID(ID)
}
