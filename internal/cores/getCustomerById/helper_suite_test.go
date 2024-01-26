package getCustomerById_test

import (
	"Test-CF/internal/cores/getCustomerById"
	"github.com/stretchr/testify/suite"
)

type getCustomerByIdExecuteTestSuite struct {
	suite.Suite

	sut getCustomerById.Service

	given  *given
	when   *when
	expect *expect
}

func (suite *getCustomerByIdExecuteTestSuite) Init() {
	suite.given = NewGiven()
	suite.when = NewWhen(suite.given)
	suite.expect = NewExpect(suite.given, suite.when)
	suite.sut = NewSut(suite.when)
}

func NewSut(context *when) getCustomerById.Service {
	return getCustomerById.NewService(
		context.getCustomerByIdPort,
	)
}
