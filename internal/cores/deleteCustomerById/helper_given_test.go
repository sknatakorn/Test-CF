package deleteCustomerById_test

import (
	"Test-CF/internal/cores/deleteCustomerById"
)

type given struct {
	request deleteCustomerById.Request
}

func NewGiven() *given {
	return &given{}
}

func (given *given) validRequest() *given {
	given.request = deleteCustomerById.Request{
		Id: 1,
	}
	return given
}
