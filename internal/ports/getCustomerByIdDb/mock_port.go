package getCustomerByIdDb

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (adaptor *Mock) Execute(id int) (*Response, error) {
	args := adaptor.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1).(error)
	}
	return args.Get(0).(*Response), nil
}
