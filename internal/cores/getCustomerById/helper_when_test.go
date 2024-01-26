package getCustomerById_test

import (
	"Test-CF/internal/ports/getCustomerByIdDb"
	"github.com/stretchr/testify/mock"
)

type when struct {
	given               *given
	getCustomerByIdDb   *getCustomerByIdDb.Response
	getCustomerByIdPort *getCustomerByIdDb.Mock
}

func NewWhen(given *given) *when {
	return &when{
		given:               given,
		getCustomerByIdDb:   &getCustomerByIdDb.Response{},
		getCustomerByIdPort: getCustomerByIdDb.NewMock(),
	}
}

func (when *when) getCustomerByIdDbOk() *when {
	when.getCustomerByIdDb = &getCustomerByIdDb.Response{
		Id:   1,
		Name: "John Doe",
		Age:  20,
	}
	when.getCustomerByIdPort.On("Execute", mock.Anything).Return(when.getCustomerByIdDb, nil)
	return when
}

func (when *when) getCustomerByIdDbError(err error) *when {
	when.getCustomerByIdPort.On("Execute", mock.Anything).Return(nil, err)
	return when
}
