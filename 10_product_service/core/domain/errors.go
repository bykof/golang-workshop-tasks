package domain

import (
	"errors"
	"fmt"
)

func ErrCouldNotFindProduct(id int64) error {
	return fmt.Errorf("could not find product with id: %d", id)
}

var (
	ErrEmptyProductName error = errors.New("the product empty cannot be empty")
)
