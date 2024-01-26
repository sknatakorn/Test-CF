package deleteCustomerById_test

import (
	"Test-CF/internal/cores/deleteCustomerById"
	"github.com/stretchr/testify/suite"
)

type deleteCustomerByIdExecuteTestSuite struct {
	suite.Suite

	sut deleteCustomerById.Service

	given  *given
	when   *when
	expect *expect
}

func (suite *deleteCustomerByIdExecuteTestSuite) Init() {
	suite.given = NewGiven()
	suite.when = NewWhen(suite.given)
	suite.expect = NewExpect(suite.given, suite.when)
	suite.sut = NewSut(suite.when)
}

func NewSut(context *when) deleteCustomerById.Service {
	return deleteCustomerById.NewService(
		context.deleteCustomerByIdPort,
	)
}
