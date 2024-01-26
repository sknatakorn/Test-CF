package createCustomer

import (
	"Test-CF/internal/ports/createCustomerDb"
)

type service struct {
	createCustomerDb createCustomerDb.Port
}

func NewService(createCustomerDb createCustomerDb.Port) Service {
	return &service{createCustomerDb: createCustomerDb}
}

func (s service) Execute(req Request) error {
	err := s.createCustomerDb.Execute(buildRequest(req))

	if err != nil {
		return err
	}
	return nil
}

func buildRequest(req Request) *createCustomerDb.Request {
	return &createCustomerDb.Request{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}
}
