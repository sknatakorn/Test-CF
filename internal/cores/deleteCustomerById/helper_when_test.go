package deleteCustomerById_test

import (
	"Test-CF/internal/ports/deleteCustomerByIdDb"
	"github.com/stretchr/testify/mock"
)

type when struct {
	given                  *given
	deleteCustomerByIdDb   *deleteCustomerByIdDb.Response
	deleteCustomerByIdPort *deleteCustomerByIdDb.Mock
}

func NewWhen(given *given) *when {
	return &when{
		given:                  given,
		deleteCustomerByIdDb:   &deleteCustomerByIdDb.Response{},
		deleteCustomerByIdPort: deleteCustomerByIdDb.NewMock(),
	}
}

func (when *when) deleteCustomerByIdDbOk() *when {
	when.deleteCustomerByIdDb = nil
	when.deleteCustomerByIdPort.On("Execute", mock.Anything).Return(nil)
	return when
}

func (when *when) deleteCustomerByIdDbError(err error) *when {
	when.deleteCustomerByIdPort.On("Execute", mock.Anything).Return(err)
	return when
}
