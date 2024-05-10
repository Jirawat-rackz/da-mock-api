package dahandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jirawat-rackz/da-mock-api/data/dadata"
	"github.com/jirawat-rackz/da-mock-api/pkg/darepository"
	"github.com/jirawat-rackz/da-mock-api/pkg/model"
)

// IDAHandler is an interface for all handler
type IDAHandler interface {
	// GetCoinList is a function to get coin list
	GetCoinList(c *gin.Context)

	// GetCoinByID is a function to get coin by id
	GetCoinByID(c *gin.Context)
}

// DAHandler is a struct for all handler
type DAHandler struct {
	DARepository darepository.IDARepository
}

// NewDAHandler is a function to create new DAHandler
func NewDAHandler(daData dadata.IDAData) IDAHandler {
	return &DAHandler{
		DARepository: darepository.NewDARepository(daData),
	}
}

// GetCoinList is a function to get coin list
func (d DAHandler) GetCoinList(c *gin.Context) {
	pagination := model.Pagination{}
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !pagination.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid pagination",
		})
		return
	}

	coins, err := d.DARepository.GetCoinList(pagination)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coins)
}

// GetCoinByID is a function to get coin by id
func (d DAHandler) GetCoinByID(c *gin.Context) {
	coinId := c.Param("id")
	coin, err := d.DARepository.GetCoinByID(coinId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, coin)
}
