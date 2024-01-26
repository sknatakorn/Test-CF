package getCustomerById_test

import (
	"Test-CF/internal/cores/getCustomerById"
)

type given struct {
	request getCustomerById.Request
}

func NewGiven() *given {
	return &given{}
}

func (given *given) validRequest() *given {
	given.request = getCustomerById.Request{
		Id: 1,
	}
	return given
}
