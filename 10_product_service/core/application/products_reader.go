package application

import (
	"encoding/json"
	"errors"
	"os"
	"product_service/core/domain"
)

func InitProductFile(jsonFilePath string) error {
	file, err := os.Create(jsonFilePath)
	if err != nil {
		return err
	}
	_, err = file.WriteString("[]")
	if err != nil {
		return err
	}
	return nil
}

func ReadProducts(jsonFilePath string) (domain.Products, error) {
	var products domain.Products
	if _, err := os.Stat(jsonFilePath); errors.Is(err, os.ErrNotExist) {
		err := InitProductFile(jsonFilePath)
		if err != nil {
			return nil, err
		}
	}

	productFile, err := os.OpenFile(jsonFilePath, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		return nil, err
	}
	defer productFile.Close()

	decoder := json.NewDecoder(productFile)
	err = decoder.Decode(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
