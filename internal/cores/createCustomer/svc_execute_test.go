package createCustomer_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestCreateCustomerExecuteTestSuite(t *testing.T) {
	suite.Run(t, new(createCustomerExecuteTestSuite))
}

func (suite *createCustomerExecuteTestSuite) SetupTest() {
	suite.Init()
}

func (suite *createCustomerExecuteTestSuite) TestExecute_ValidRequest_AllGreen() {
	suite.given.
		validRequest()
	suite.when.createCustomerByDbOk()

	actualErr := suite.sut.Execute(suite.given.request)

	assert.Equal(suite.T(), nil, actualErr)
}

func (suite *createCustomerExecuteTestSuite) TestExecute_ValidRequest_CreateCustomerError() {
	suite.given.
		validRequest()
	suite.when.createCustomerDbError(errors.New(""))
	actualErr := suite.sut.Execute(suite.given.request)
	assert.Equal(suite.T(), "", actualErr.Error())
}
