package updateCustomerById_test

import (
	"Test-CF/internal/cores/updateCustomerById"
	"github.com/stretchr/testify/suite"
)

type updateCustomerByIdExecuteTestSuite struct {
	suite.Suite

	sut updateCustomerById.Service

	given  *given
	when   *when
	expect *expect
}

func (suite *updateCustomerByIdExecuteTestSuite) Init() {
	suite.given = NewGiven()
	suite.when = NewWhen(suite.given)
	suite.expect = NewExpect(suite.given, suite.when)
	suite.sut = NewSut(suite.when)
}

func NewSut(context *when) updateCustomerById.Service {
	return updateCustomerById.NewService(
		context.updateCustomerByIdPort,
	)
}
