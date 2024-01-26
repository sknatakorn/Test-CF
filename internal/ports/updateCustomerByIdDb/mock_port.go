package updateCustomerByIdDb

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (adaptor *Mock) Execute(request *Request) error {
	args := adaptor.Called(request)
	if args.Error(0) != nil {
		return args.Error(0).(error)
	}
	return nil
}
