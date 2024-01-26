package getCustomerById_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestGetCustomerByIdExecuteTestSuite(t *testing.T) {
	suite.Run(t, new(getCustomerByIdExecuteTestSuite))
}

func (suite *getCustomerByIdExecuteTestSuite) SetupTest() {
	suite.Init()
}

func (suite *getCustomerByIdExecuteTestSuite) TestExecute_ValidRequest_AllGreen() {
	suite.given.
		validRequest()
	suite.when.getCustomerByIdDbOk()

	actualResp, actualErr := suite.sut.Execute(suite.given.request)

	assert.Equal(suite.T(), suite.expect.responseOk(), actualResp)
	assert.Equal(suite.T(), nil, actualErr)
}

func (suite *getCustomerByIdExecuteTestSuite) TestExecute_ValidRequest_GetCustomerByIdError() {
	suite.given.
		validRequest()
	suite.when.getCustomerByIdDbError(errors.New(""))
	_, actualErr := suite.sut.Execute(suite.given.request)
	assert.Equal(suite.T(), "", actualErr.Error())
}
