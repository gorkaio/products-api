package main

import (
	"errors"
	"time"
	"unicode"
	"unicode/utf8"
)

type ProductId string
type Timestamp time.Time

type Product struct {
	Id ProductId
	Price int
	Description string
	CreatedAt Timestamp
}

func NewProduct(id ProductId, description string, price int) (Product, error) {
	id, err := validateId(id)
	if (err != nil) {
		return Product{}, err
	}

	description, err = validateDescription(description)
	if (err != nil) {
		return Product{}, err
	}

	return Product{
		Id: ProductId(id),
		Price: price,
		CreatedAt: Timestamp(time.Now()),
		Description: description,
	}, nil
}

func validateId(id ProductId) (ProductId, error) {
	if isASCII(string(id)) {
		return id, nil
	}

	return "", errors.New("invalid id")
}

func validateDescription(description string) (string, error) {
	if (!utf8.ValidString(description) || utf8.RuneCountInString(description) > 50) {
		return "", errors.New("invalid description")
	}

	return description, nil
}

func isASCII(s string) bool {
    for _, c := range s {
        if c > unicode.MaxASCII {
            return false
        }
    }

    return true
}