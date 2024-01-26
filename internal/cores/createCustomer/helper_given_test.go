package createCustomer_test

import (
	"Test-CF/internal/cores/createCustomer"
)

type given struct {
	request createCustomer.Request
}

func NewGiven() *given {
	return &given{}
}

func (given *given) validRequest() *given {
	given.request = createCustomer.Request{
		Id:   1,
		Name: "John Doe",
		Age:  20,
	}
	return given
}
