package updateCustomerById

import (
	"Test-CF/internal/ports/updateCustomerByIdDb"
)

type service struct {
	updateCustomerByIdDb updateCustomerByIdDb.Port
}

func NewService(updateCustomerByIdDb updateCustomerByIdDb.Port) Service {
	return &service{updateCustomerByIdDb: updateCustomerByIdDb}
}

func (s service) Execute(req Request) error {
	err := s.updateCustomerByIdDb.Execute(buildRequest(req))

	if err != nil {
		return err
	}
	return nil
}

func buildRequest(req Request) *updateCustomerByIdDb.Request {
	return &updateCustomerByIdDb.Request{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}
}
