package updateCustomerById_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestUpdateCustomerByIdExecuteTestSuite(t *testing.T) {
	suite.Run(t, new(updateCustomerByIdExecuteTestSuite))
}

func (suite *updateCustomerByIdExecuteTestSuite) SetupTest() {
	suite.Init()
}

func (suite *updateCustomerByIdExecuteTestSuite) TestExecute_ValidRequest_AllGreen() {
	suite.given.
		validRequest()
	suite.when.updateCustomerByIdOk()

	actualErr := suite.sut.Execute(suite.given.request)

	assert.Equal(suite.T(), nil, actualErr)
}

func (suite *updateCustomerByIdExecuteTestSuite) TestExecute_ValidRequest_UpdateCustomerByIdError() {
	suite.given.
		validRequest()
	suite.when.updateCustomerByIdError(errors.New(""))
	actualErr := suite.sut.Execute(suite.given.request)
	assert.Equal(suite.T(), "", actualErr.Error())
}
