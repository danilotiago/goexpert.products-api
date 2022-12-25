package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	expectedName := "Notebook"
	expectedPrice := 4000.0
	product, err := NewProduct(expectedName, expectedPrice)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, expectedName, product.Name)
	assert.Equal(t, expectedPrice, product.Price)
}

func TestProduct_whenEmptyName(t *testing.T) {
	expectedName := ""
	expectedPrice := 4000.0

	product, err := NewProduct(expectedName, expectedPrice)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProduct_whenPriceIsZero(t *testing.T) {
	expectedName := "Notebook"
	expectedPrice := 0.0

	product, err := NewProduct(expectedName, expectedPrice)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProduct_whenPriceIsNegative(t *testing.T) {
	expectedName := "Notebook"
	expectedPrice := -10.0

	product, err := NewProduct(expectedName, expectedPrice)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
