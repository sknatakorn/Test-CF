package getCustomerById_test

import (
	"Test-CF/internal/cores/getCustomerById"
)

type expect struct {
	given *given
	when  *when
}

func NewExpect(given *given, when *when) *expect {
	return &expect{
		given: given,
		when:  when,
	}
}

func (expect *expect) responseOk() *getCustomerById.Response {
	return &getCustomerById.Response{
		Id:   expect.when.getCustomerByIdDb.Id,
		Name: expect.when.getCustomerByIdDb.Name,
		Age:  expect.when.getCustomerByIdDb.Age,
	}
}
