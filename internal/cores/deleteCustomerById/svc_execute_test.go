package deleteCustomerById_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDeleteCustomerByIdExecuteTestSuite(t *testing.T) {
	suite.Run(t, new(deleteCustomerByIdExecuteTestSuite))
}

func (suite *deleteCustomerByIdExecuteTestSuite) SetupTest() {
	suite.Init()
}

func (suite *deleteCustomerByIdExecuteTestSuite) TestExecute_ValidRequest_AllGreen() {
	suite.given.
		validRequest()
	suite.when.deleteCustomerByIdDbOk()

	actualErr := suite.sut.Execute(suite.given.request)

	assert.Equal(suite.T(), nil, actualErr)
}

func (suite *deleteCustomerByIdExecuteTestSuite) TestExecute_ValidRequest_DeleteCustomerByIdError() {
	suite.given.
		validRequest()
	suite.when.deleteCustomerByIdDbError(errors.New(""))
	actualErr := suite.sut.Execute(suite.given.request)
	assert.Equal(suite.T(), "", actualErr.Error())
}
