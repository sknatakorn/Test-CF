package updateCustomerById_test

import (
	updateCustomerByIdDb "Test-CF/internal/ports/updateCustomerByIdDb"
	"github.com/stretchr/testify/mock"
)

type when struct {
	given                  *given
	updateCustomerByIdDb   *updateCustomerByIdDb.Response
	updateCustomerByIdPort *updateCustomerByIdDb.Mock
}

func NewWhen(given *given) *when {
	return &when{
		given:                  given,
		updateCustomerByIdDb:   &updateCustomerByIdDb.Response{},
		updateCustomerByIdPort: updateCustomerByIdDb.NewMock(),
	}
}

func (when *when) updateCustomerByIdOk() *when {
	when.updateCustomerByIdDb = nil
	when.updateCustomerByIdPort.On("Execute", mock.Anything).Return(nil)
	return when
}

func (when *when) updateCustomerByIdError(err error) *when {
	when.updateCustomerByIdPort.On("Execute", mock.Anything).Return(err)
	return when
}
