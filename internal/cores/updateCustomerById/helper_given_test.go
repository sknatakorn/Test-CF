package updateCustomerById_test

import (
	"Test-CF/internal/cores/updateCustomerById"
)

type given struct {
	request updateCustomerById.Request
}

func NewGiven() *given {
	return &given{}
}

func (given *given) validRequest() *given {
	given.request = updateCustomerById.Request{
		Id:   1,
		Name: "John Doe",
		Age:  20,
	}
	return given
}
