package getCustomerById

import (
	"Test-CF/internal/ports/getCustomerByIdDb"
)

type service struct {
	getCustomerByIdDb getCustomerByIdDb.Port
}

func NewService(getCustomerByIdDb getCustomerByIdDb.Port) Service {
	return &service{getCustomerByIdDb: getCustomerByIdDb}
}

func (s service) Execute(req Request) (*Response, error) {
	customer, err := s.getCustomerByIdDb.Execute(req.Id)

	if err != nil {
		return nil, err
	}
	return buildResponse(customer), nil
}

func buildResponse(res *getCustomerByIdDb.Response) *Response {
	return &Response{
		Id:   res.Id,
		Name: res.Name,
		Age:  res.Age,
	}
}
