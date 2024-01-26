package deleteCustomerById

import (
	"Test-CF/internal/ports/deleteCustomerByIdDb"
)

type service struct {
	deleteCustomerByIdDb deleteCustomerByIdDb.Port
}

func NewService(deleteCustomerByIdDb deleteCustomerByIdDb.Port) Service {
	return &service{deleteCustomerByIdDb: deleteCustomerByIdDb}
}

func (s service) Execute(req Request) error {
	err := s.deleteCustomerByIdDb.Execute(buildRequest(req))

	if err != nil {
		return err
	}
	return nil
}

func buildRequest(request Request) *deleteCustomerByIdDb.Request {
	return &deleteCustomerByIdDb.Request{Id: request.Id}
}
