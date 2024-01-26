package createCustomer_test

import (
	"Test-CF/internal/cores/createCustomer"
	"github.com/stretchr/testify/suite"
)

type createCustomerExecuteTestSuite struct {
	suite.Suite

	sut createCustomer.Service

	given  *given
	when   *when
	expect *expect
}

func (suite *createCustomerExecuteTestSuite) Init() {
	suite.given = NewGiven()
	suite.when = NewWhen(suite.given)
	suite.expect = NewExpect(suite.given, suite.when)
	suite.sut = NewSut(suite.when)
}

func NewSut(context *when) createCustomer.Service {
	return createCustomer.NewService(
		context.createCustomerPort,
	)
}
