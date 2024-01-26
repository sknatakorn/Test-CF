package createCustomer_test

import (
	"Test-CF/internal/ports/createCustomerDb"
	"github.com/stretchr/testify/mock"
)

type when struct {
	given              *given
	createCustomerDb   *createCustomerDb.Response
	createCustomerPort *createCustomerDb.Mock
}

func NewWhen(given *given) *when {
	return &when{
		given:              given,
		createCustomerDb:   &createCustomerDb.Response{},
		createCustomerPort: createCustomerDb.NewMock(),
	}
}

func (when *when) createCustomerByDbOk() *when {
	when.createCustomerDb = nil
	when.createCustomerPort.On("Execute", mock.Anything).Return(nil)
	return when
}

func (when *when) createCustomerDbError(err error) *when {
	when.createCustomerPort.On("Execute", mock.Anything).Return(err)
	return when
}
